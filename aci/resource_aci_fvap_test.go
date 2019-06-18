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

func TestAccAciApplicationProfile_Basic(t *testing.T) {
	var application_profile models.ApplicationProfile
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	description := "application_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciApplicationProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciApplicationProfileConfig_basic(fv_tenant_name, fv_ap_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciApplicationProfileExists("aci_application_profile.fooapplication_profile", &application_profile),
					testAccCheckAciApplicationProfileAttributes(fv_tenant_name, fv_ap_name, description, &application_profile),
				),
			},
		},
	})
}

func testAccCheckAciApplicationProfileConfig_basic(fv_tenant_name, fv_ap_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_application_profile" "fooapplication_profile" {
		name 		= "%s"
		description = "application_profile created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, fv_ap_name)
}

func testAccCheckAciApplicationProfileExists(name string, application_profile *models.ApplicationProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Application Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Application Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		application_profileFound := models.ApplicationProfileFromContainer(cont)
		if application_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Application Profile %s not found", rs.Primary.ID)
		}
		*application_profile = *application_profileFound
		return nil
	}
}

func testAccCheckAciApplicationProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_application_profile" {
			cont, err := client.Get(rs.Primary.ID)
			application_profile := models.ApplicationProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Application Profile %s Still exists", application_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciApplicationProfileAttributes(fv_tenant_name, fv_ap_name, description string, application_profile *models.ApplicationProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ap_name != GetMOName(application_profile.DistinguishedName) {
			return fmt.Errorf("Bad fv_ap %s", GetMOName(application_profile.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(application_profile.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(application_profile.DistinguishedName)))
		}
		if description != application_profile.Description {
			return fmt.Errorf("Bad application_profile Description %s", application_profile.Description)
		}

		return nil
	}
}
