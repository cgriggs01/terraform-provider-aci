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

func TestAccAciL3Outside_Basic(t *testing.T) {
	var l3_outside models.L3Outside
	fv_tenant_name := acctest.RandString(5)
	l3ext_out_name := acctest.RandString(5)
	description := "l3_outside created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciL3OutsideDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciL3OutsideConfig_basic(fv_tenant_name, l3ext_out_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciL3OutsideExists("aci_l3_outside.fool3_outside", &l3_outside),
					testAccCheckAciL3OutsideAttributes(fv_tenant_name, l3ext_out_name, description, &l3_outside),
				),
			},
		},
	})
}

func testAccCheckAciL3OutsideConfig_basic(fv_tenant_name, l3ext_out_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_l3_outside" "fool3_outside" {
		name 		= "%s"
		description = "l3_outside created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, l3ext_out_name)
}

func testAccCheckAciL3OutsideExists(name string, l3_outside *models.L3Outside) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("L3 Outside %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No L3 Outside dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		l3_outsideFound := models.L3OutsideFromContainer(cont)
		if l3_outsideFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("L3 Outside %s not found", rs.Primary.ID)
		}
		*l3_outside = *l3_outsideFound
		return nil
	}
}

func testAccCheckAciL3OutsideDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_l3_outside" {
			cont, err := client.Get(rs.Primary.ID)
			l3_outside := models.L3OutsideFromContainer(cont)
			if err == nil {
				return fmt.Errorf("L3 Outside %s Still exists", l3_outside.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciL3OutsideAttributes(fv_tenant_name, l3ext_out_name, description string, l3_outside *models.L3Outside) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if l3ext_out_name != GetMOName(l3_outside.DistinguishedName) {
			return fmt.Errorf("Bad l3ext_out %s", GetMOName(l3_outside.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(l3_outside.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(l3_outside.DistinguishedName)))
		}
		if description != l3_outside.Description {
			return fmt.Errorf("Bad l3_outside Description %s", l3_outside.Description)
		}

		return nil
	}
}
