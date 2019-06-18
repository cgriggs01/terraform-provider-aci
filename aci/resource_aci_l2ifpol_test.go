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

func TestAccAciL2InterfacePolicy_Basic(t *testing.T) {
	var l2_interface_policy models.L2InterfacePolicy
	l2_if_pol_name := acctest.RandString(5)
	description := "l2_interface_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciL2InterfacePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciL2InterfacePolicyConfig_basic(l2_if_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciL2InterfacePolicyExists("aci_l2_interface_policy.fool2_interface_policy", &l2_interface_policy),
					testAccCheckAciL2InterfacePolicyAttributes(l2_if_pol_name, description, &l2_interface_policy),
				),
			},
		},
	})
}

func testAccCheckAciL2InterfacePolicyConfig_basic(l2_if_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_l2_interface_policy" "fool2_interface_policy" {
		name 		= "%s"
		description = "l2_interface_policy created while acceptance testing"

	}

	`, l2_if_pol_name)
}

func testAccCheckAciL2InterfacePolicyExists(name string, l2_interface_policy *models.L2InterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("L2 Interface Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No L2 Interface Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		l2_interface_policyFound := models.L2InterfacePolicyFromContainer(cont)
		if l2_interface_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("L2 Interface Policy %s not found", rs.Primary.ID)
		}
		*l2_interface_policy = *l2_interface_policyFound
		return nil
	}
}

func testAccCheckAciL2InterfacePolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_l2_interface_policy" {
			cont, err := client.Get(rs.Primary.ID)
			l2_interface_policy := models.L2InterfacePolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("L2 Interface Policy %s Still exists", l2_interface_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciL2InterfacePolicyAttributes(l2_if_pol_name, description string, l2_interface_policy *models.L2InterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if l2_if_pol_name != GetMOName(l2_interface_policy.DistinguishedName) {
			return fmt.Errorf("Bad l2_if_pol %s", GetMOName(l2_interface_policy.DistinguishedName))
		}

		if description != l2_interface_policy.Description {
			return fmt.Errorf("Bad l2_interface_policy Description %s", l2_interface_policy.Description)
		}

		return nil
	}
}
