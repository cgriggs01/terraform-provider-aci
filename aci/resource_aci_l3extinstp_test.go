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

func TestAccAciExternalNetworkInstanceProfile_Basic(t *testing.T) {
	var external_network_instance_profile models.ExternalNetworkInstanceProfile
	fv_tenant_name := acctest.RandString(5)
	l3ext_out_name := acctest.RandString(5)
	l3ext_inst_p_name := acctest.RandString(5)
	description := "external_network_instance_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciExternalNetworkInstanceProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciExternalNetworkInstanceProfileConfig_basic(fv_tenant_name, l3ext_out_name, l3ext_inst_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciExternalNetworkInstanceProfileExists("aci_external_network_instance_profile.fooexternal_network_instance_profile", &external_network_instance_profile),
					testAccCheckAciExternalNetworkInstanceProfileAttributes(fv_tenant_name, l3ext_out_name, l3ext_inst_p_name, description, &external_network_instance_profile),
				),
			},
		},
	})
}

func testAccCheckAciExternalNetworkInstanceProfileConfig_basic(fv_tenant_name, l3ext_out_name, l3ext_inst_p_name string) string {
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

	resource "aci_external_network_instance_profile" "fooexternal_network_instance_profile" {
		name 		= "%s"
		description = "external_network_instance_profile created while acceptance testing"
		l3_outside_dn = "${aci_l3_outside.fool3_outside.id}"
	}

	`, fv_tenant_name, l3ext_out_name, l3ext_inst_p_name)
}

func testAccCheckAciExternalNetworkInstanceProfileExists(name string, external_network_instance_profile *models.ExternalNetworkInstanceProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("External Network Instance Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No External Network Instance Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		external_network_instance_profileFound := models.ExternalNetworkInstanceProfileFromContainer(cont)
		if external_network_instance_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("External Network Instance Profile %s not found", rs.Primary.ID)
		}
		*external_network_instance_profile = *external_network_instance_profileFound
		return nil
	}
}

func testAccCheckAciExternalNetworkInstanceProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_external_network_instance_profile" {
			cont, err := client.Get(rs.Primary.ID)
			external_network_instance_profile := models.ExternalNetworkInstanceProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("External Network Instance Profile %s Still exists", external_network_instance_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciExternalNetworkInstanceProfileAttributes(fv_tenant_name, l3ext_out_name, l3ext_inst_p_name, description string, external_network_instance_profile *models.ExternalNetworkInstanceProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if l3ext_inst_p_name != GetMOName(external_network_instance_profile.DistinguishedName) {
			return fmt.Errorf("Bad l3ext_inst_p %s", GetMOName(external_network_instance_profile.DistinguishedName))
		}

		if l3ext_out_name != GetMOName(GetParentDn(external_network_instance_profile.DistinguishedName)) {
			return fmt.Errorf(" Bad l3ext_out %s", GetMOName(GetParentDn(external_network_instance_profile.DistinguishedName)))
		}
		if description != external_network_instance_profile.Description {
			return fmt.Errorf("Bad external_network_instance_profile Description %s", external_network_instance_profile.Description)
		}

		return nil
	}
}
