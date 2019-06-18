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

func TestAccAciSubnet_Basic(t *testing.T) {
	var subnet models.Subnet
	fv_tenant_name := acctest.RandString(5)
	fv_bd_name := acctest.RandString(5)
	fv_subnet_name := acctest.RandString(5)
	description := "subnet created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciSubnetConfig_basic(fv_tenant_name, fv_bd_name, fv_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciSubnetExists("aci_subnet.foosubnet", &subnet),
					testAccCheckAciSubnetAttributes(fv_tenant_name, fv_bd_name, fv_subnet_name, description, &subnet),
				),
			},
		},
	})
}

func testAccCheckAciSubnetConfig_basic(fv_tenant_name, fv_bd_name, fv_subnet_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_bridge_domain" "foobridge_domain" {
		name 		= "%s"
		description = "bridge_domain created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_subnet" "foosubnet" {
		name 		= "%s"
		description = "subnet created while acceptance testing"
		bridge_domain_dn = "${aci_bridge_domain.foobridge_domain.id}"
	}

	`, fv_tenant_name, fv_bd_name, fv_subnet_name)
}

func testAccCheckAciSubnetExists(name string, subnet *models.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Subnet %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Subnet dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		subnetFound := models.SubnetFromContainer(cont)
		if subnetFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Subnet %s not found", rs.Primary.ID)
		}
		*subnet = *subnetFound
		return nil
	}
}

func testAccCheckAciSubnetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_subnet" {
			cont, err := client.Get(rs.Primary.ID)
			subnet := models.SubnetFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Subnet %s Still exists", subnet.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciSubnetAttributes(fv_tenant_name, fv_bd_name, fv_subnet_name, description string, subnet *models.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_subnet_name != GetMOName(subnet.DistinguishedName) {
			return fmt.Errorf("Bad fv_subnet %s", GetMOName(subnet.DistinguishedName))
		}

		if fv_bd_name != GetMOName(GetParentDn(subnet.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_bd %s", GetMOName(GetParentDn(subnet.DistinguishedName)))
		}
		if description != subnet.Description {
			return fmt.Errorf("Bad subnet Description %s", subnet.Description)
		}

		return nil
	}
}
