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

func TestAccAciAny_Basic(t *testing.T) {
	var any models.Any
	fv_tenant_name := acctest.RandString(5)
	fv_ctx_name := acctest.RandString(5)
	vz_any_name := acctest.RandString(5)
	description := "any created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAnyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAnyConfig_basic(fv_tenant_name, fv_ctx_name, vz_any_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAnyExists("aci_any.fooany", &any),
					testAccCheckAciAnyAttributes(fv_tenant_name, fv_ctx_name, vz_any_name, description, &any),
				),
			},
		},
	})
}

func testAccCheckAciAnyConfig_basic(fv_tenant_name, fv_ctx_name, vz_any_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_vrf" "foovrf" {
		name 		= "%s"
		description = "vrf created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_any" "fooany" {
		name 		= "%s"
		description = "any created while acceptance testing"
		vrf_dn = "${aci_vrf.foovrf.id}"
	}

	`, fv_tenant_name, fv_ctx_name, vz_any_name)
}

func testAccCheckAciAnyExists(name string, any *models.Any) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Any %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Any dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		anyFound := models.AnyFromContainer(cont)
		if anyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Any %s not found", rs.Primary.ID)
		}
		*any = *anyFound
		return nil
	}
}

func testAccCheckAciAnyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_any" {
			cont, err := client.Get(rs.Primary.ID)
			any := models.AnyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Any %s Still exists", any.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAnyAttributes(fv_tenant_name, fv_ctx_name, vz_any_name, description string, any *models.Any) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_any_name != GetMOName(any.DistinguishedName) {
			return fmt.Errorf("Bad vz_any %s", GetMOName(any.DistinguishedName))
		}

		if fv_ctx_name != GetMOName(GetParentDn(any.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_ctx %s", GetMOName(GetParentDn(any.DistinguishedName)))
		}
		if description != any.Description {
			return fmt.Errorf("Bad any Description %s", any.Description)
		}

		return nil
	}
}
