package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAciInterfaceFCPolicy_Basic(t *testing.T) {
	var interface_fc_policy models.InterfaceFCPolicy
	fc_if_pol_name := acctest.RandString(5)
	description := "interface_fc_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciInterfaceFCPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciInterfaceFCPolicyConfig_basic(fc_if_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciInterfaceFCPolicyExists("aci_interface_fc_policy.foointerface_fc_policy", &interface_fc_policy),
					testAccCheckAciInterfaceFCPolicyAttributes(fc_if_pol_name, description, &interface_fc_policy),
				),
			},
		},
	})
}

func testAccCheckAciInterfaceFCPolicyConfig_basic(fc_if_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_interface_fc_policy" "foointerface_fc_policy" {
		name 		= "%s"
		description = "interface_fc_policy created while acceptance testing"

	}

	`, fc_if_pol_name)
}

func testAccCheckAciInterfaceFCPolicyExists(name string, interface_fc_policy *models.InterfaceFCPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Interface FC Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Interface FC Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		interface_fc_policyFound := models.InterfaceFCPolicyFromContainer(cont)
		if interface_fc_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Interface FC Policy %s not found", rs.Primary.ID)
		}
		*interface_fc_policy = *interface_fc_policyFound
		return nil
	}
}

func testAccCheckAciInterfaceFCPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_interface_fc_policy" {
			cont, err := client.Get(rs.Primary.ID)
			interface_fc_policy := models.InterfaceFCPolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Interface FC Policy %s Still exists", interface_fc_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciInterfaceFCPolicyAttributes(fc_if_pol_name, description string, interface_fc_policy *models.InterfaceFCPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fc_if_pol_name != GetMOName(interface_fc_policy.DistinguishedName) {
			return fmt.Errorf("Bad fc_if_pol %s", GetMOName(interface_fc_policy.DistinguishedName))
		}

		if description != interface_fc_policy.Description {
			return fmt.Errorf("Bad interface_fc_policy Description %s", interface_fc_policy.Description)
		}

		return nil
	}
}
