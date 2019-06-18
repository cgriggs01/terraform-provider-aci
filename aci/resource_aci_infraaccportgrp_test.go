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

func TestAccAciLeafAccessPortPolicyGroup_Basic(t *testing.T) {
	var leaf_access_port_policy_group models.LeafAccessPortPolicyGroup
	infra_acc_port_grp_name := acctest.RandString(5)
	description := "leaf_access_port_policy_group created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciLeafAccessPortPolicyGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciLeafAccessPortPolicyGroupConfig_basic(infra_acc_port_grp_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciLeafAccessPortPolicyGroupExists("aci_leaf_access_port_policy_group.fooleaf_access_port_policy_group", &leaf_access_port_policy_group),
					testAccCheckAciLeafAccessPortPolicyGroupAttributes(infra_acc_port_grp_name, description, &leaf_access_port_policy_group),
				),
			},
		},
	})
}

func testAccCheckAciLeafAccessPortPolicyGroupConfig_basic(infra_acc_port_grp_name string) string {
	return fmt.Sprintf(`

	resource "aci_leaf_access_port_policy_group" "fooleaf_access_port_policy_group" {
		name 		= "%s"
		description = "leaf_access_port_policy_group created while acceptance testing"

	}

	`, infra_acc_port_grp_name)
}

func testAccCheckAciLeafAccessPortPolicyGroupExists(name string, leaf_access_port_policy_group *models.LeafAccessPortPolicyGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Leaf Access Port Policy Group %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Leaf Access Port Policy Group dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		leaf_access_port_policy_groupFound := models.LeafAccessPortPolicyGroupFromContainer(cont)
		if leaf_access_port_policy_groupFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Leaf Access Port Policy Group %s not found", rs.Primary.ID)
		}
		*leaf_access_port_policy_group = *leaf_access_port_policy_groupFound
		return nil
	}
}

func testAccCheckAciLeafAccessPortPolicyGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_leaf_access_port_policy_group" {
			cont, err := client.Get(rs.Primary.ID)
			leaf_access_port_policy_group := models.LeafAccessPortPolicyGroupFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Leaf Access Port Policy Group %s Still exists", leaf_access_port_policy_group.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciLeafAccessPortPolicyGroupAttributes(infra_acc_port_grp_name, description string, leaf_access_port_policy_group *models.LeafAccessPortPolicyGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_acc_port_grp_name != GetMOName(leaf_access_port_policy_group.DistinguishedName) {
			return fmt.Errorf("Bad infra_acc_port_grp %s", GetMOName(leaf_access_port_policy_group.DistinguishedName))
		}

		if description != leaf_access_port_policy_group.Description {
			return fmt.Errorf("Bad leaf_access_port_policy_group Description %s", leaf_access_port_policy_group.Description)
		}

		return nil
	}
}
