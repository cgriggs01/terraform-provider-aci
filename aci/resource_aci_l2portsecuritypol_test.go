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

func TestAccAciPortSecurityPolicy_Basic(t *testing.T) {
	var port_security_policy models.PortSecurityPolicy
	l2_port_security_pol_name := acctest.RandString(5)
	description := "port_security_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciPortSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciPortSecurityPolicyConfig_basic(l2_port_security_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciPortSecurityPolicyExists("aci_port_security_policy.fooport_security_policy", &port_security_policy),
					testAccCheckAciPortSecurityPolicyAttributes(l2_port_security_pol_name, description, &port_security_policy),
				),
			},
		},
	})
}

func testAccCheckAciPortSecurityPolicyConfig_basic(l2_port_security_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_port_security_policy" "fooport_security_policy" {
		name 		= "%s"
		description = "port_security_policy created while acceptance testing"

	}

	`, l2_port_security_pol_name)
}

func testAccCheckAciPortSecurityPolicyExists(name string, port_security_policy *models.PortSecurityPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Port Security Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Port Security Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		port_security_policyFound := models.PortSecurityPolicyFromContainer(cont)
		if port_security_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Port Security Policy %s not found", rs.Primary.ID)
		}
		*port_security_policy = *port_security_policyFound
		return nil
	}
}

func testAccCheckAciPortSecurityPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_port_security_policy" {
			cont, err := client.Get(rs.Primary.ID)
			port_security_policy := models.PortSecurityPolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Port Security Policy %s Still exists", port_security_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciPortSecurityPolicyAttributes(l2_port_security_pol_name, description string, port_security_policy *models.PortSecurityPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if l2_port_security_pol_name != GetMOName(port_security_policy.DistinguishedName) {
			return fmt.Errorf("Bad l2_port_security_pol %s", GetMOName(port_security_policy.DistinguishedName))
		}

		if description != port_security_policy.Description {
			return fmt.Errorf("Bad port_security_policy Description %s", port_security_policy.Description)
		}

		return nil
	}
}
