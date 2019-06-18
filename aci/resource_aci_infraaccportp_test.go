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

func TestAccAciLeafInterfaceProfile_Basic(t *testing.T) {
	var leaf_interface_profile models.LeafInterfaceProfile
	infra_acc_port_p_name := acctest.RandString(5)
	description := "leaf_interface_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciLeafInterfaceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciLeafInterfaceProfileConfig_basic(infra_acc_port_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciLeafInterfaceProfileExists("aci_leaf_interface_profile.fooleaf_interface_profile", &leaf_interface_profile),
					testAccCheckAciLeafInterfaceProfileAttributes(infra_acc_port_p_name, description, &leaf_interface_profile),
				),
			},
		},
	})
}

func testAccCheckAciLeafInterfaceProfileConfig_basic(infra_acc_port_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_leaf_interface_profile" "fooleaf_interface_profile" {
		name 		= "%s"
		description = "leaf_interface_profile created while acceptance testing"

	}

	`, infra_acc_port_p_name)
}

func testAccCheckAciLeafInterfaceProfileExists(name string, leaf_interface_profile *models.LeafInterfaceProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Leaf Interface Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Leaf Interface Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		leaf_interface_profileFound := models.LeafInterfaceProfileFromContainer(cont)
		if leaf_interface_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Leaf Interface Profile %s not found", rs.Primary.ID)
		}
		*leaf_interface_profile = *leaf_interface_profileFound
		return nil
	}
}

func testAccCheckAciLeafInterfaceProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_leaf_interface_profile" {
			cont, err := client.Get(rs.Primary.ID)
			leaf_interface_profile := models.LeafInterfaceProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Leaf Interface Profile %s Still exists", leaf_interface_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciLeafInterfaceProfileAttributes(infra_acc_port_p_name, description string, leaf_interface_profile *models.LeafInterfaceProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_acc_port_p_name != GetMOName(leaf_interface_profile.DistinguishedName) {
			return fmt.Errorf("Bad infra_acc_port_p %s", GetMOName(leaf_interface_profile.DistinguishedName))
		}

		if description != leaf_interface_profile.Description {
			return fmt.Errorf("Bad leaf_interface_profile Description %s", leaf_interface_profile.Description)
		}

		return nil
	}
}
