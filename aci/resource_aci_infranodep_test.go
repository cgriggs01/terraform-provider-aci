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

func TestAccAciLeafProfile_Basic(t *testing.T) {
	var leaf_profile models.LeafProfile
	infra_node_p_name := acctest.RandString(5)
	description := "leaf_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciLeafProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciLeafProfileConfig_basic(infra_node_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciLeafProfileExists("aci_leaf_profile.fooleaf_profile", &leaf_profile),
					testAccCheckAciLeafProfileAttributes(infra_node_p_name, description, &leaf_profile),
				),
			},
		},
	})
}

func testAccCheckAciLeafProfileConfig_basic(infra_node_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_leaf_profile" "fooleaf_profile" {
		name 		= "%s"
		description = "leaf_profile created while acceptance testing"

	}

	`, infra_node_p_name)
}

func testAccCheckAciLeafProfileExists(name string, leaf_profile *models.LeafProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Leaf Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Leaf Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		leaf_profileFound := models.LeafProfileFromContainer(cont)
		if leaf_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Leaf Profile %s not found", rs.Primary.ID)
		}
		*leaf_profile = *leaf_profileFound
		return nil
	}
}

func testAccCheckAciLeafProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_leaf_profile" {
			cont, err := client.Get(rs.Primary.ID)
			leaf_profile := models.LeafProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Leaf Profile %s Still exists", leaf_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciLeafProfileAttributes(infra_node_p_name, description string, leaf_profile *models.LeafProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_node_p_name != GetMOName(leaf_profile.DistinguishedName) {
			return fmt.Errorf("Bad infra_node_p %s", GetMOName(leaf_profile.DistinguishedName))
		}

		if description != leaf_profile.Description {
			return fmt.Errorf("Bad leaf_profile Description %s", leaf_profile.Description)
		}

		return nil
	}
}
