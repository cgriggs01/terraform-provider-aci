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

func TestAccAciAttachableAccessEntityProfile_Basic(t *testing.T) {
	var attachable_access_entity_profile models.AttachableAccessEntityProfile
	infra_att_entity_p_name := acctest.RandString(5)
	description := "attachable_access_entity_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAttachableAccessEntityProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAttachableAccessEntityProfileConfig_basic(infra_att_entity_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAttachableAccessEntityProfileExists("aci_attachable_access_entity_profile.fooattachable_access_entity_profile", &attachable_access_entity_profile),
					testAccCheckAciAttachableAccessEntityProfileAttributes(infra_att_entity_p_name, description, &attachable_access_entity_profile),
				),
			},
		},
	})
}

func testAccCheckAciAttachableAccessEntityProfileConfig_basic(infra_att_entity_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_attachable_access_entity_profile" "fooattachable_access_entity_profile" {
		name 		= "%s"
		description = "attachable_access_entity_profile created while acceptance testing"

	}

	`, infra_att_entity_p_name)
}

func testAccCheckAciAttachableAccessEntityProfileExists(name string, attachable_access_entity_profile *models.AttachableAccessEntityProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Attachable Access Entity Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Attachable Access Entity Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		attachable_access_entity_profileFound := models.AttachableAccessEntityProfileFromContainer(cont)
		if attachable_access_entity_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Attachable Access Entity Profile %s not found", rs.Primary.ID)
		}
		*attachable_access_entity_profile = *attachable_access_entity_profileFound
		return nil
	}
}

func testAccCheckAciAttachableAccessEntityProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_attachable_access_entity_profile" {
			cont, err := client.Get(rs.Primary.ID)
			attachable_access_entity_profile := models.AttachableAccessEntityProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Attachable Access Entity Profile %s Still exists", attachable_access_entity_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAttachableAccessEntityProfileAttributes(infra_att_entity_p_name, description string, attachable_access_entity_profile *models.AttachableAccessEntityProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_att_entity_p_name != GetMOName(attachable_access_entity_profile.DistinguishedName) {
			return fmt.Errorf("Bad infra_att_entity_p %s", GetMOName(attachable_access_entity_profile.DistinguishedName))
		}

		if description != attachable_access_entity_profile.Description {
			return fmt.Errorf("Bad attachable_access_entity_profile Description %s", attachable_access_entity_profile.Description)
		}

		return nil
	}
}
