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

func TestAccAciAccessPortBlock_Basic(t *testing.T) {
	var access_port_block models.AccessPortBlock
	infra_acc_port_p_name := acctest.RandString(5)
	infra_h_port_s_name := acctest.RandString(5)
	infra_port_blk_name := acctest.RandString(5)
	description := "access_port_block created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAccessPortBlockDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAccessPortBlockConfig_basic(infra_acc_port_p_name, infra_h_port_s_name, infra_port_blk_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAccessPortBlockExists("aci_access_port_block.fooaccess_port_block", &access_port_block),
					testAccCheckAciAccessPortBlockAttributes(infra_acc_port_p_name, infra_h_port_s_name, infra_port_blk_name, description, &access_port_block),
				),
			},
		},
	})
}

func testAccCheckAciAccessPortBlockConfig_basic(infra_acc_port_p_name, infra_h_port_s_name, infra_port_blk_name string) string {
	return fmt.Sprintf(`

	resource "aci_leaf_interface_profile" "fooleaf_interface_profile" {
		name 		= "%s"
		description = "leaf_interface_profile created while acceptance testing"

	}

	resource "aci_access_port_selector" "fooaccess_port_selector" {
		name 		= "%s"
		description = "access_port_selector created while acceptance testing"
		leaf_interface_profile_dn = "${aci_leaf_interface_profile.fooleaf_interface_profile.id}"
	}

	resource "aci_access_port_block" "fooaccess_port_block" {
		name 		= "%s"
		description = "access_port_block created while acceptance testing"
		access_port_selector_dn = "${aci_access_port_selector.fooaccess_port_selector.id}"
	}

	`, infra_acc_port_p_name, infra_h_port_s_name, infra_port_blk_name)
}

func testAccCheckAciAccessPortBlockExists(name string, access_port_block *models.AccessPortBlock) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Access Port Block %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Access Port Block dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		access_port_blockFound := models.AccessPortBlockFromContainer(cont)
		if access_port_blockFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Access Port Block %s not found", rs.Primary.ID)
		}
		*access_port_block = *access_port_blockFound
		return nil
	}
}

func testAccCheckAciAccessPortBlockDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_access_port_block" {
			cont, err := client.Get(rs.Primary.ID)
			access_port_block := models.AccessPortBlockFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Access Port Block %s Still exists", access_port_block.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAccessPortBlockAttributes(infra_acc_port_p_name, infra_h_port_s_name, infra_port_blk_name, description string, access_port_block *models.AccessPortBlock) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_port_blk_name != GetMOName(access_port_block.DistinguishedName) {
			return fmt.Errorf("Bad infra_port_blk %s", GetMOName(access_port_block.DistinguishedName))
		}

		if infra_h_port_s_name != GetMOName(GetParentDn(access_port_block.DistinguishedName)) {
			return fmt.Errorf(" Bad infra_h_port_s %s", GetMOName(GetParentDn(access_port_block.DistinguishedName)))
		}
		if description != access_port_block.Description {
			return fmt.Errorf("Bad access_port_block Description %s", access_port_block.Description)
		}

		return nil
	}
}
