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

func TestAccAciVlanEncapsulationforVxlanTraffic_Basic(t *testing.T) {
	var vlan_encapsulationfor_vxlan_traffic models.VlanEncapsulationforVxlanTraffic
	infra_att_entity_p_name := acctest.RandString(5)
	infra_prov_acc_name := acctest.RandString(5)
	description := "vlan_encapsulationfor_vxlan_traffic created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciVlanEncapsulationforVxlanTrafficDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciVlanEncapsulationforVxlanTrafficConfig_basic(infra_att_entity_p_name, infra_prov_acc_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciVlanEncapsulationforVxlanTrafficExists("aci_vlan_encapsulationfor_vxlan_traffic.foovlan_encapsulationfor_vxlan_traffic", &vlan_encapsulationfor_vxlan_traffic),
					testAccCheckAciVlanEncapsulationforVxlanTrafficAttributes(infra_att_entity_p_name, infra_prov_acc_name, description, &vlan_encapsulationfor_vxlan_traffic),
				),
			},
		},
	})
}

func testAccCheckAciVlanEncapsulationforVxlanTrafficConfig_basic(infra_att_entity_p_name, infra_prov_acc_name string) string {
	return fmt.Sprintf(`

	resource "aci_attachable_access_entity_profile" "fooattachable_access_entity_profile" {
		name 		= "%s"
		description = "attachable_access_entity_profile created while acceptance testing"

	}

	resource "aci_vlan_encapsulationfor_vxlan_traffic" "foovlan_encapsulationfor_vxlan_traffic" {
		name 		= "%s"
		description = "vlan_encapsulationfor_vxlan_traffic created while acceptance testing"
		attachable_access_entity_profile_dn = "${aci_attachable_access_entity_profile.fooattachable_access_entity_profile.id}"
	}

	`, infra_att_entity_p_name, infra_prov_acc_name)
}

func testAccCheckAciVlanEncapsulationforVxlanTrafficExists(name string, vlan_encapsulationfor_vxlan_traffic *models.VlanEncapsulationforVxlanTraffic) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Vlan Encapsulation for Vxlan Traffic %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vlan Encapsulation for Vxlan Traffic dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		vlan_encapsulationfor_vxlan_trafficFound := models.VlanEncapsulationforVxlanTrafficFromContainer(cont)
		if vlan_encapsulationfor_vxlan_trafficFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Vlan Encapsulation for Vxlan Traffic %s not found", rs.Primary.ID)
		}
		*vlan_encapsulationfor_vxlan_traffic = *vlan_encapsulationfor_vxlan_trafficFound
		return nil
	}
}

func testAccCheckAciVlanEncapsulationforVxlanTrafficDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_vlan_encapsulationfor_vxlan_traffic" {
			cont, err := client.Get(rs.Primary.ID)
			vlan_encapsulationfor_vxlan_traffic := models.VlanEncapsulationforVxlanTrafficFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Vlan Encapsulation for Vxlan Traffic %s Still exists", vlan_encapsulationfor_vxlan_traffic.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciVlanEncapsulationforVxlanTrafficAttributes(infra_att_entity_p_name, infra_prov_acc_name, description string, vlan_encapsulationfor_vxlan_traffic *models.VlanEncapsulationforVxlanTraffic) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_prov_acc_name != GetMOName(vlan_encapsulationfor_vxlan_traffic.DistinguishedName) {
			return fmt.Errorf("Bad infra_prov_acc %s", GetMOName(vlan_encapsulationfor_vxlan_traffic.DistinguishedName))
		}

		if infra_att_entity_p_name != GetMOName(GetParentDn(vlan_encapsulationfor_vxlan_traffic.DistinguishedName)) {
			return fmt.Errorf(" Bad infra_att_entity_p %s", GetMOName(GetParentDn(vlan_encapsulationfor_vxlan_traffic.DistinguishedName)))
		}
		if description != vlan_encapsulationfor_vxlan_traffic.Description {
			return fmt.Errorf("Bad vlan_encapsulationfor_vxlan_traffic Description %s", vlan_encapsulationfor_vxlan_traffic.Description)
		}

		return nil
	}
}
