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

func TestAccAciLACPPolicy_Basic(t *testing.T) {
	var lacp_policy models.LACPPolicy
	lacp_lag_pol_name := acctest.RandString(5)
	description := "lacp_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciLACPPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciLACPPolicyConfig_basic(lacp_lag_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciLACPPolicyExists("aci_lacp_policy.foolacp_policy", &lacp_policy),
					testAccCheckAciLACPPolicyAttributes(lacp_lag_pol_name, description, &lacp_policy),
				),
			},
		},
	})
}

func testAccCheckAciLACPPolicyConfig_basic(lacp_lag_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_lacp_policy" "foolacp_policy" {
		name 		= "%s"
		description = "lacp_policy created while acceptance testing"

	}

	`, lacp_lag_pol_name)
}

func testAccCheckAciLACPPolicyExists(name string, lacp_policy *models.LACPPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("LACP Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LACP Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		lacp_policyFound := models.LACPPolicyFromContainer(cont)
		if lacp_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("LACP Policy %s not found", rs.Primary.ID)
		}
		*lacp_policy = *lacp_policyFound
		return nil
	}
}

func testAccCheckAciLACPPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_lacp_policy" {
			cont, err := client.Get(rs.Primary.ID)
			lacp_policy := models.LACPPolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("LACP Policy %s Still exists", lacp_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciLACPPolicyAttributes(lacp_lag_pol_name, description string, lacp_policy *models.LACPPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if lacp_lag_pol_name != GetMOName(lacp_policy.DistinguishedName) {
			return fmt.Errorf("Bad lacp_lag_pol %s", GetMOName(lacp_policy.DistinguishedName))
		}

		if description != lacp_policy.Description {
			return fmt.Errorf("Bad lacp_policy Description %s", lacp_policy.Description)
		}

		return nil
	}
}
