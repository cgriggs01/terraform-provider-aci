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

func TestAccAciAccessPortSelector_Basic(t *testing.T) {
	var access_port_selector models.AccessPortSelector
	infra_acc_port_p_name := acctest.RandString(5)
	infra_h_port_s_name := acctest.RandString(5)
	description := "access_port_selector created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAccessPortSelectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAccessPortSelectorConfig_basic(infra_acc_port_p_name, infra_h_port_s_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAccessPortSelectorExists("aci_access_port_selector.fooaccess_port_selector", &access_port_selector),
					testAccCheckAciAccessPortSelectorAttributes(infra_acc_port_p_name, infra_h_port_s_name, description, &access_port_selector),
				),
			},
		},
	})
}

func testAccCheckAciAccessPortSelectorConfig_basic(infra_acc_port_p_name, infra_h_port_s_name string) string {
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

	`, infra_acc_port_p_name, infra_h_port_s_name)
}

func testAccCheckAciAccessPortSelectorExists(name string, access_port_selector *models.AccessPortSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Access Port Selector %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Access Port Selector dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		access_port_selectorFound := models.AccessPortSelectorFromContainer(cont)
		if access_port_selectorFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Access Port Selector %s not found", rs.Primary.ID)
		}
		*access_port_selector = *access_port_selectorFound
		return nil
	}
}

func testAccCheckAciAccessPortSelectorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_access_port_selector" {
			cont, err := client.Get(rs.Primary.ID)
			access_port_selector := models.AccessPortSelectorFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Access Port Selector %s Still exists", access_port_selector.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAccessPortSelectorAttributes(infra_acc_port_p_name, infra_h_port_s_name, description string, access_port_selector *models.AccessPortSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_h_port_s_name != GetMOName(access_port_selector.DistinguishedName) {
			return fmt.Errorf("Bad infra_h_port_s %s", GetMOName(access_port_selector.DistinguishedName))
		}

		if infra_acc_port_p_name != GetMOName(GetParentDn(access_port_selector.DistinguishedName)) {
			return fmt.Errorf(" Bad infra_acc_port_p %s", GetMOName(GetParentDn(access_port_selector.DistinguishedName)))
		}
		if description != access_port_selector.Description {
			return fmt.Errorf("Bad access_port_selector Description %s", access_port_selector.Description)
		}

		return nil
	}
}
