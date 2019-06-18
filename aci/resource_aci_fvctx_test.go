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

func TestAccAciVRF_Basic(t *testing.T) {
	var vrf models.VRF
	fv_tenant_name := acctest.RandString(5)
	fv_ctx_name := acctest.RandString(5)
	description := "vrf created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciVRFDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciVRFConfig_basic(fv_tenant_name, fv_ctx_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciVRFExists("aci_vrf.foovrf", &vrf),
					testAccCheckAciVRFAttributes(fv_tenant_name, fv_ctx_name, description, &vrf),
				),
			},
		},
	})
}

func testAccCheckAciVRFConfig_basic(fv_tenant_name, fv_ctx_name string) string {
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

	`, fv_tenant_name, fv_ctx_name)
}

func testAccCheckAciVRFExists(name string, vrf *models.VRF) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VRF %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VRF dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		vrfFound := models.VRFFromContainer(cont)
		if vrfFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("VRF %s not found", rs.Primary.ID)
		}
		*vrf = *vrfFound
		return nil
	}
}

func testAccCheckAciVRFDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_vrf" {
			cont, err := client.Get(rs.Primary.ID)
			vrf := models.VRFFromContainer(cont)
			if err == nil {
				return fmt.Errorf("VRF %s Still exists", vrf.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciVRFAttributes(fv_tenant_name, fv_ctx_name, description string, vrf *models.VRF) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ctx_name != GetMOName(vrf.DistinguishedName) {
			return fmt.Errorf("Bad fv_ctx %s", GetMOName(vrf.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(vrf.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(vrf.DistinguishedName)))
		}
		if description != vrf.Description {
			return fmt.Errorf("Bad vrf Description %s", vrf.Description)
		}

		return nil
	}
}
