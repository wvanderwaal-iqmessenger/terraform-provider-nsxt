/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vmware/go-vmware-nsxt"
	"net/http"
	"testing"
)

func TestNSXL4PortNsServiceBasic(t *testing.T) {

	serviceName := fmt.Sprintf("test-nsx-l4-service")
	updateServiceName := fmt.Sprintf("%s-update", serviceName)
	testResourceName := "nsxt_l4_port_set_ns_service.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNSXL4ServiceCheckDestroy(state, serviceName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNSXserviceCreateTemplate(serviceName, "TCP", "99"),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXL4ServiceExists(serviceName, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", serviceName),
					resource.TestCheckResourceAttr(testResourceName, "description", "l4 service"),
					resource.TestCheckResourceAttr(testResourceName, "l4_protocol", "TCP"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccNSXserviceCreateTemplate(updateServiceName, "UDP", "98"),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXL4ServiceExists(updateServiceName, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", updateServiceName),
					resource.TestCheckResourceAttr(testResourceName, "description", "l4 service"),
					resource.TestCheckResourceAttr(testResourceName, "l4_protocol", "UDP"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "1"),
				),
			},
		},
	})
}

func testAccNSXL4ServiceExists(display_name string, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		nsxClient := testAccProvider.Meta().(*nsxt.APIClient)

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("NSX L4 ns service resource %s not found in resources", resourceName)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("NSX L4 ns service resource ID not set in resources ")
		}

		service, responseCode, err := nsxClient.GroupingObjectsApi.ReadL4PortSetNSService(nsxClient.Context, resourceID)
		if err != nil {
			return fmt.Errorf("Error while retrieving L4 ns service ID %s. Error: %v", resourceID, err)
		}

		if responseCode.StatusCode != http.StatusOK {
			return fmt.Errorf("Error while checking if L4 ns service %s exists. HTTP return code was %d", resourceID, responseCode)
		}

		if display_name == service.DisplayName {
			return nil
		}
		return fmt.Errorf("NSX L4 ns service %s wasn't found", display_name)
	}
}

func testAccNSXL4ServiceCheckDestroy(state *terraform.State, display_name string) error {

	nsxClient := testAccProvider.Meta().(*nsxt.APIClient)

	for _, rs := range state.RootModule().Resources {

		if rs.Type != "nsxt_l4_port_set_ns_service" {
			continue
		}

		resourceID := rs.Primary.Attributes["id"]
		service, responseCode, err := nsxClient.GroupingObjectsApi.ReadL4PortSetNSService(nsxClient.Context, resourceID)
		if err != nil {
			if responseCode.StatusCode != http.StatusOK {
				return nil
			}
			return fmt.Errorf("Error while retrieving L4 ns service ID %s. Error: %v", resourceID, err)
		}

		if display_name == service.DisplayName {
			return fmt.Errorf("NSX L4 ns service %s still exists", display_name)
		}
	}
	return nil
}

func testAccNSXserviceCreateTemplate(serviceName string, protocol string, port string) string {
	return fmt.Sprintf(`
resource "nsxt_l4_port_set_ns_service" "test" {
    description = "l4 service"
    display_name = "%s"
    l4_protocol = "%s"
    destination_ports = [ "%s" ]
    tags = [{scope = "scope1"
             tag = "tag1"}
    ]
}`, serviceName, protocol, port)
}