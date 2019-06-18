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

func TestAccAciMiscablingProtocolInterfacePolicy_Basic(t *testing.T) {
	var miscabling_protocol_interface_policy models.MiscablingProtocolInterfacePolicy
	mcp_if_pol_name := acctest.RandString(5)
	description := "mis-cabling_protocol_interface_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciMiscablingProtocolInterfacePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciMiscablingProtocolInterfacePolicyConfig_basic(mcp_if_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciMiscablingProtocolInterfacePolicyExists("aci_mis-cabling_protocol_interface_policy.foomis-cabling_protocol_interface_policy", &miscabling_protocol_interface_policy),
					testAccCheckAciMiscablingProtocolInterfacePolicyAttributes(mcp_if_pol_name, description, &miscabling_protocol_interface_policy),
				),
			},
		},
	})
}

func testAccCheckAciMiscablingProtocolInterfacePolicyConfig_basic(mcp_if_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_mis-cabling_protocol_interface_policy" "foomis-cabling_protocol_interface_policy" {
		name 		= "%s"
		description = "mis-cabling_protocol_interface_policy created while acceptance testing"

	}

	`, mcp_if_pol_name)
}

func testAccCheckAciMiscablingProtocolInterfacePolicyExists(name string, miscabling_protocol_interface_policy *models.MiscablingProtocolInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Mis-cabling Protocol Interface Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Mis-cabling Protocol Interface Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		miscabling_protocol_interface_policyFound := models.MiscablingProtocolInterfacePolicyFromContainer(cont)
		if miscabling_protocol_interface_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Mis-cabling Protocol Interface Policy %s not found", rs.Primary.ID)
		}
		*miscabling_protocol_interface_policy = *miscabling_protocol_interface_policyFound
		return nil
	}
}

func testAccCheckAciMiscablingProtocolInterfacePolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_mis-cabling_protocol_interface_policy" {
			cont, err := client.Get(rs.Primary.ID)
			miscabling_protocol_interface_policy := models.MiscablingProtocolInterfacePolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Mis-cabling Protocol Interface Policy %s Still exists", miscabling_protocol_interface_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciMiscablingProtocolInterfacePolicyAttributes(mcp_if_pol_name, description string, miscabling_protocol_interface_policy *models.MiscablingProtocolInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if mcp_if_pol_name != GetMOName(miscabling_protocol_interface_policy.DistinguishedName) {
			return fmt.Errorf("Bad mcp_if_pol %s", GetMOName(miscabling_protocol_interface_policy.DistinguishedName))
		}

		if description != miscabling_protocol_interface_policy.Description {
			return fmt.Errorf("Bad miscabling_protocol_interface_policy Description %s", miscabling_protocol_interface_policy.Description)
		}

		return nil
	}
}
