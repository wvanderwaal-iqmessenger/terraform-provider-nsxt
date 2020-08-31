/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	gm_domains "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/domains"
	gm_model "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/domains"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

func resourceNsxtPolicyFixedGatewayPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtPolicyFixedGatewayPolicyCreate,
		Read:   resourceNsxtPolicyFixedGatewayPolicyRead,
		Update: resourceNsxtPolicyFixedGatewayPolicyUpdate,
		Delete: resourceNsxtPolicyFixedGatewayPolicyDelete,

		Schema: getPolicyFixedGatewayPolicySchema(),
	}
}

func getPolicyFixedGatewayPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"path":         getPolicyPathSchema(true, true, "Path for this Gateway Policy"),
		"description":  getComputedDescriptionSchema(),
		"tag":          getTagsSchema(),
		"rule":         getSecurityPolicyAndGatewayRulesSchema(false),
		"default_rule": getPolicyDefaultRulesSchema(),
		"revision":     getRevisionSchema(),
	}
}

func getPolicyDefaultRulesSchema() *schema.Schema {
	return &schema.Schema{
		Type:          schema.TypeList,
		Description:   "List of default rules",
		Optional:      true,
		ConflictsWith: []string{"rule"},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"nsx_id":      getComputedNsxIDSchema(),
				"scope":       getPolicyPathSchema(true, false, "Scope for this rule"),
				"description": getComputedDescriptionSchema(),
				"path":        getPathSchema(),
				"revision":    getRevisionSchema(),
				"logged": {
					Type:        schema.TypeBool,
					Description: "Flag to enable packet logging",
					Optional:    true,
					Default:     false,
				},
				"tag": getTagsSchema(),
				"log_label": {
					Type:        schema.TypeString,
					Description: "Additional information (string) which will be propagated to the rule syslog",
					Optional:    true,
				},
				"action": {
					Type:         schema.TypeString,
					Description:  "Action",
					Optional:     true,
					ValidateFunc: validation.StringInSlice(securityPolicyActionValues, false),
					Default:      model.Rule_ACTION_ALLOW,
				},
				"sequence_number": {
					Type:        schema.TypeInt,
					Description: "Sequence number of the this rule",
					Computed:    true,
				},
			},
		},
	}
}

func updateGatewayPolicyDefaultRuleByScope(rule model.Rule, d *schema.ResourceData, connector *client.RestConnector, isGlobalManager bool) *model.Rule {
	defaultRules := d.Get("default_rule").([]interface{})

	for _, obj := range defaultRules {
		defaultRule := obj.(map[string]interface{})
		scope := defaultRule["scope"].(string)

		if len(rule.Scope) == 1 && scope == rule.Scope[0] {
			description := defaultRule["description"].(string)
			rule.Description = &description
			action := defaultRule["action"].(string)
			rule.Action = &action
			logLabel := defaultRule["log_label"].(string)
			rule.Tag = &logLabel
			logged := defaultRule["logged"].(bool)
			rule.Logged = &logged
			tags := getPolicyTagsFromSet(defaultRule["tag"].(*schema.Set))
			if len(tags) > 0 || len(rule.Tags) > 0 {
				rule.Tags = tags
			}

			log.Printf("[DEBUG] Updating Default Rule with ID %s", *rule.Id)
			return &rule
		}
	}

	// This rule is not present in new config - check if was just deleted
	// If so, the rule needs to be reverted
	_, oldRules := d.GetChange("default_rule")
	for _, oldRule := range oldRules.([]interface{}) {
		oldRuleMap := oldRule.(map[string]interface{})
		if oldID, ok := oldRuleMap["nsx_id"]; ok {
			if (rule.Id != nil) && (*rule.Id == oldID.(string)) {
				rule, err := revertDefaultRuleByScope(rule, connector, isGlobalManager)
				if err != nil {
					log.Printf("[WARNING]: Failed to revert rule: %s", err)
				}

				log.Printf("[DEBUG] Reverting Default Rule with ID %s", *rule.Id)
				return &rule
			}
		}
	}

	return nil
}

func setPolicyDefaultRulesInSchema(d *schema.ResourceData, rules []model.Rule) error {
	var rulesList []map[string]interface{}
	for _, rule := range rules {
		elem := make(map[string]interface{})
		elem["description"] = rule.Description
		elem["log_label"] = rule.Tag
		elem["logged"] = rule.Logged
		elem["action"] = rule.Action
		elem["revision"] = rule.Revision
		setPathListInMap(elem, "scope", rule.Scope)
		elem["sequence_number"] = rule.SequenceNumber
		elem["tag"] = initPolicyTagsSet(rule.Tags)
		elem["path"] = rule.Path
		elem["nsx_id"] = rule.Id

		rulesList = append(rulesList, elem)
	}

	return d.Set("default_rule", rulesList)
}

func revertPolicyFixedGatewayPolicy(fixedPolicy model.GatewayPolicy, m interface{}) (model.GatewayPolicy, error) {
	connector := getPolicyConnector(m)
	isGlobalManager := isPolicyGlobalManager(m)

	// Default values for Name and Description are ID
	fixedPolicy.Rules = nil
	fixedPolicy.Description = fixedPolicy.DisplayName

	var childRules []*data.StructValue

	for _, rule := range fixedPolicy.Rules {
		if rule.IsDefault != nil && *rule.IsDefault {
			revertedRule, err := revertDefaultRuleByScope(rule, connector, isGlobalManager)
			if err != nil {
				return model.GatewayPolicy{}, fmt.Errorf("[WARNING]: Failed to revert default rule: %s", err)
			}
			childRule, err := createPolicyChildRule(*revertedRule.Id, revertedRule, false)
			if err != nil {
				return model.GatewayPolicy{}, err
			}
			childRules = append(childRules, childRule)
		}
	}

	if len(childRules) > 0 {
		fixedPolicy.Children = childRules
	}

	if len(fixedPolicy.Tags) > 0 {
		tags := make([]model.Tag, 0)
		fixedPolicy.Tags = tags
	}

	return fixedPolicy, nil
}

func revertDefaultRuleByScope(rule model.Rule, connector *client.RestConnector, isGlobalManager bool) (model.Rule, error) {
	if len(rule.Scope) != 1 {
		return model.Rule{}, fmt.Errorf("Expected default rule %s to have single scope", *rule.Path)
	}

	if len(rule.Tags) > 0 {
		tags := make([]model.Tag, 0)
		rule.Tags = tags
	}

	if !strings.Contains(rule.Scope[0], "infra/tier-0s") {
		// This rule is not based on T0
		defaultAction := "DROP"
		rule.DisplayName = &rule.Scope[0]
		rule.Description = &rule.Scope[0]
		rule.Action = &defaultAction
		return rule, nil
	}

	// Default rule default values are set according to scope Tier0
	gwID := getPolicyIDFromPath(rule.Scope[0])
	gw, err := getPolicyTier0Gateway(gwID, connector, isGlobalManager)
	if err != nil {
		return rule, fmt.Errorf("Failed to retrieve scope object %s for rule %s", gwID, *rule.Path)
	}

	rule.DisplayName = gw.DisplayName
	rule.Description = gw.Description
	defaultAction := "DROP"
	if gw.ForceWhitelisting != nil && *gw.ForceWhitelisting {
		defaultAction = "ALLOW"
	}
	rule.Action = &defaultAction

	return rule, nil
}

func createPolicyChildRule(ruleID string, rule model.Rule, shouldDelete bool) (*data.StructValue, error) {
	converter := bindings.NewTypeConverter()
	converter.SetMode(bindings.REST)

	childRule := model.ChildRule{
		ResourceType: "ChildRule",
		//Id:              &ruleID,
		Rule:            &rule,
		MarkedForDelete: &shouldDelete,
	}

	dataValue, errors := converter.ConvertToVapi(childRule, model.ChildRuleBindingType())
	if len(errors) > 0 {
		return nil, errors[0]
	}

	return dataValue.(*data.StructValue), nil
}

func createChildDomainWithGatewayPolicy(domain string, policyID string, policy model.GatewayPolicy) (*data.StructValue, error) {
	converter := bindings.NewTypeConverter()
	converter.SetMode(bindings.REST)

	childPolicy := model.ChildGatewayPolicy{
		//Id:            &policyID,
		ResourceType:  "ChildGatewayPolicy",
		GatewayPolicy: &policy,
	}

	dataValue, errors := converter.ConvertToVapi(childPolicy, model.ChildGatewayPolicyBindingType())
	if len(errors) > 0 {
		return nil, errors[0]
	}

	var domainChildren []*data.StructValue
	domainChildren = append(domainChildren, dataValue.(*data.StructValue))

	targetType := "Domain"
	childDomain := model.ChildResourceReference{
		Id:           &domain,
		ResourceType: "ChildResourceReference",
		TargetType:   &targetType,
		Children:     domainChildren,
	}

	dataValue, errors = converter.ConvertToVapi(childDomain, model.ChildResourceReferenceBindingType())
	if len(errors) > 0 {
		return nil, errors[0]
	}
	return dataValue.(*data.StructValue), nil
}

func fixedPolicyInfraPatch(policy model.GatewayPolicy, domain string, m interface{}) error {
	childDomain, err := createChildDomainWithGatewayPolicy(domain, *policy.Id, policy)
	if err != nil {
		return fmt.Errorf("Failed to create H-API for Fixed Gateway Policy: %s", err)
	}

	var infraChildren []*data.StructValue
	infraChildren = append(infraChildren, childDomain)

	infraType := "Infra"
	infraObj := model.Infra{
		Children:     infraChildren,
		ResourceType: &infraType,
	}

	return policyInfraPatch(infraObj, isPolicyGlobalManager(m), getPolicyConnector(m), true)

}

func updatePolicyFixedGatewayPolicy(id string, d *schema.ResourceData, m interface{}) error {

	connector := getPolicyConnector(m)
	isGlobalManager := isPolicyGlobalManager(m)
	path := d.Get("path").(string)
	domain := getDomainFromResourcePath(path)

	if domain == "" {
		return fmt.Errorf("Failed to extract domain from Gateway Policy path %s", path)
	}

	fixedPolicy, err := getGatewayPolicyInDomain(id, domain, connector, isPolicyGlobalManager(m))
	if err != nil {
		return err
	}

	fixedPolicy.Rules = nil
	if d.HasChange("description") {
		description := d.Get("description").(string)
		fixedPolicy.Description = &description
	}

	if d.HasChange("tag") {
		fixedPolicy.Tags = getPolicyTagsFromSchema(d)
	}

	var childRules []*data.StructValue
	if d.HasChange("rule") {
		oldRules, _ := d.GetChange("rule")
		rules := getPolicyRulesFromSchema(d)

		existingRules := make(map[string]bool)
		for _, rule := range rules {
			ruleID := newUUID()
			if rule.Id != nil {
				ruleID = *rule.Id
				existingRules[ruleID] = true
			} else {
				rule.Id = &ruleID
			}

			childRule, err := createPolicyChildRule(ruleID, rule, false)
			if err != nil {
				return err
			}
			log.Printf("[DEBUG]: Adding child rule with id %s", ruleID)
			childRules = append(childRules, childRule)
		}

		for _, oldRule := range oldRules.([]interface{}) {
			oldRuleMap := oldRule.(map[string]interface{})
			oldRuleID := oldRuleMap["nsx_id"].(string)
			if _, exists := existingRules[oldRuleID]; !exists {
				resourceType := "Rule"
				rule := model.Rule{
					Id:           &oldRuleID,
					ResourceType: &resourceType,
				}

				childRule, err := createPolicyChildRule(oldRuleID, rule, true)
				if err != nil {
					return err
				}
				log.Printf("[DEBUG]: Deleting child rule with id %s", oldRuleID)
				childRules = append(childRules, childRule)

			}
		}
	}

	if d.HasChange("default_rule") {
		for _, existingDefaultRule := range fixedPolicy.Rules {
			if existingDefaultRule.IsDefault != nil && *existingDefaultRule.IsDefault {
				updatedDefaultRule := updateGatewayPolicyDefaultRuleByScope(existingDefaultRule, d, connector, isGlobalManager)
				if updatedDefaultRule != nil {
					childRule, err := createPolicyChildRule(*updatedDefaultRule.Id, *updatedDefaultRule, false)
					if err != nil {
						return err
					}
					childRules = append(childRules, childRule)
				}
			}
		}
	}

	log.Printf("[DEBUG]: Updating default policy %s with %d child rules", id, len(childRules))
	if len(childRules) > 0 {
		fixedPolicy.Children = childRules
	}

	err = fixedPolicyInfraPatch(fixedPolicy, domain, m)
	if err != nil {
		return handleUpdateError("Fixed Gateway Policy", id, err)
	}

	return nil
}

func resourceNsxtPolicyFixedGatewayPolicyCreate(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	id := getPolicyIDFromPath(path)

	if id == "" {
		return fmt.Errorf("Failed to extract ID from Gateway Policy path %s", path)
	}

	err := updatePolicyFixedGatewayPolicy(id, d, m)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceNsxtPolicyFixedGatewayPolicyRead(d, m)
}

func resourceNsxtPolicyFixedGatewayPolicyRead(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Gateway Policy ID")
	}

	path := d.Get("path").(string)
	domain := getDomainFromResourcePath(path)

	var obj model.GatewayPolicy
	if isPolicyGlobalManager(m) {
		client := gm_domains.NewDefaultGatewayPoliciesClient(connector)
		gmObj, err := client.Get(domain, id)
		if err != nil {
			return handleReadError(d, "Fixed Gateway Policy", id, err)
		}
		rawObj, err := convertModelBindingType(gmObj, gm_model.GatewayPolicyBindingType(), model.GatewayPolicyBindingType())
		if err != nil {
			return err
		}
		obj = rawObj.(model.GatewayPolicy)
	} else {
		var err error
		client := domains.NewDefaultGatewayPoliciesClient(connector)
		obj, err = client.Get(domain, id)
		if err != nil {
			return handleReadError(d, "Fixed Gateway Policy", id, err)
		}
	}

	d.Set("description", obj.Description)
	setPolicyTagsInSchema(d, obj.Tags)
	d.Set("path", obj.Path)
	d.Set("domain", getDomainFromResourcePath(*obj.Path))
	d.Set("revision", obj.Revision)

	var rules []model.Rule
	var defaultRules []model.Rule

	for _, rule := range obj.Rules {
		if rule.IsDefault != nil && *rule.IsDefault {
			defaultRules = append(defaultRules, rule)
		} else {
			rules = append(rules, rule)
		}
	}

	err := setPolicyRulesInSchema(d, rules)
	if err != nil {
		return err
	}
	return setPolicyDefaultRulesInSchema(d, defaultRules)
}

func resourceNsxtPolicyFixedGatewayPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Fixed Gateway Policy ID")
	}
	err := updatePolicyFixedGatewayPolicy(id, d, m)
	if err != nil {
		return err
	}

	return resourceNsxtPolicyFixedGatewayPolicyRead(d, m)
}

func resourceNsxtPolicyFixedGatewayPolicyDelete(d *schema.ResourceData, m interface{}) error {
	// Delete means revert back to default
	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining Fixed Gateway Policy ID")
	}

	path := d.Get("path").(string)
	domain := getDomainFromResourcePath(path)

	fixedPolicy, err := getGatewayPolicyInDomain(id, domain, getPolicyConnector(m), isPolicyGlobalManager(m))
	if err != nil {
		return err
	}

	revertedPolicy, err := revertPolicyFixedGatewayPolicy(fixedPolicy, m)
	if err != nil {
		return fmt.Errorf("Failed to revert Fixed Gateway Policy %s: %s", id, err)
	}

	err = fixedPolicyInfraPatch(revertedPolicy, domain, m)
	if err != nil {
		return handleUpdateError("Fixed Gateway Policy", id, err)
	}
	return nil
}
