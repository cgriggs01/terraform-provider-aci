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

func TestAccAciApplicationEPG_Basic(t *testing.T) {
	var application_epg models.ApplicationEPG
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	fv_ae_pg_name := acctest.RandString(5)
	description := "application_epg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciApplicationEPGDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciApplicationEPGConfig_basic(fv_tenant_name, fv_ap_name, fv_ae_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciApplicationEPGExists("aci_application_epg.fooapplication_epg", &application_epg),
					testAccCheckAciApplicationEPGAttributes(fv_tenant_name, fv_ap_name, fv_ae_pg_name, description, &application_epg),
				),
			},
		},
	})
}

func testAccCheckAciApplicationEPGConfig_basic(fv_tenant_name, fv_ap_name, fv_ae_pg_name string) string {
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

	resource "aci_application_epg" "fooapplication_epg" {
		name 		= "%s"
		description = "application_epg created while acceptance testing"
		application_profile_dn = "${aci_application_profile.fooapplication_profile.id}"
	}

	`, fv_tenant_name, fv_ap_name, fv_ae_pg_name)
}

func testAccCheckAciApplicationEPGExists(name string, application_epg *models.ApplicationEPG) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Application EPG %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Application EPG dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		application_epgFound := models.ApplicationEPGFromContainer(cont)
		if application_epgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Application EPG %s not found", rs.Primary.ID)
		}
		*application_epg = *application_epgFound
		return nil
	}
}

func testAccCheckAciApplicationEPGDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_application_epg" {
			cont, err := client.Get(rs.Primary.ID)
			application_epg := models.ApplicationEPGFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Application EPG %s Still exists", application_epg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciApplicationEPGAttributes(fv_tenant_name, fv_ap_name, fv_ae_pg_name, description string, application_epg *models.ApplicationEPG) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ae_pg_name != GetMOName(application_epg.DistinguishedName) {
			return fmt.Errorf("Bad fv_ae_pg %s", GetMOName(application_epg.DistinguishedName))
		}

		if fv_ap_name != GetMOName(GetParentDn(application_epg.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_ap %s", GetMOName(GetParentDn(application_epg.DistinguishedName)))
		}
		if description != application_epg.Description {
			return fmt.Errorf("Bad application_epg Description %s", application_epg.Description)
		}

		return nil
	}
}
