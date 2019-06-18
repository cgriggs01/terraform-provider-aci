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

func TestAccAciLLDPInterfacePolicy_Basic(t *testing.T) {
	var lldp_interface_policy models.LLDPInterfacePolicy
	lldp_if_pol_name := acctest.RandString(5)
	description := "lldp_interface_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciLLDPInterfacePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciLLDPInterfacePolicyConfig_basic(lldp_if_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciLLDPInterfacePolicyExists("aci_lldp_interface_policy.foolldp_interface_policy", &lldp_interface_policy),
					testAccCheckAciLLDPInterfacePolicyAttributes(lldp_if_pol_name, description, &lldp_interface_policy),
				),
			},
		},
	})
}

func testAccCheckAciLLDPInterfacePolicyConfig_basic(lldp_if_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_lldp_interface_policy" "foolldp_interface_policy" {
		name 		= "%s"
		description = "lldp_interface_policy created while acceptance testing"

	}

	`, lldp_if_pol_name)
}

func testAccCheckAciLLDPInterfacePolicyExists(name string, lldp_interface_policy *models.LLDPInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("LLDP Interface Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LLDP Interface Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		lldp_interface_policyFound := models.LLDPInterfacePolicyFromContainer(cont)
		if lldp_interface_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("LLDP Interface Policy %s not found", rs.Primary.ID)
		}
		*lldp_interface_policy = *lldp_interface_policyFound
		return nil
	}
}

func testAccCheckAciLLDPInterfacePolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_lldp_interface_policy" {
			cont, err := client.Get(rs.Primary.ID)
			lldp_interface_policy := models.LLDPInterfacePolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("LLDP Interface Policy %s Still exists", lldp_interface_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciLLDPInterfacePolicyAttributes(lldp_if_pol_name, description string, lldp_interface_policy *models.LLDPInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if lldp_if_pol_name != GetMOName(lldp_interface_policy.DistinguishedName) {
			return fmt.Errorf("Bad lldp_if_pol %s", GetMOName(lldp_interface_policy.DistinguishedName))
		}

		if description != lldp_interface_policy.Description {
			return fmt.Errorf("Bad lldp_interface_policy Description %s", lldp_interface_policy.Description)
		}

		return nil
	}
}
