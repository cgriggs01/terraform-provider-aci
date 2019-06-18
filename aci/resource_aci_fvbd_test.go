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

func TestAccAciBridgeDomain_Basic(t *testing.T) {
	var bridge_domain models.BridgeDomain
	fv_tenant_name := acctest.RandString(5)
	fv_bd_name := acctest.RandString(5)
	description := "bridge_domain created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciBridgeDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciBridgeDomainConfig_basic(fv_tenant_name, fv_bd_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciBridgeDomainExists("aci_bridge_domain.foobridge_domain", &bridge_domain),
					testAccCheckAciBridgeDomainAttributes(fv_tenant_name, fv_bd_name, description, &bridge_domain),
				),
			},
		},
	})
}

func testAccCheckAciBridgeDomainConfig_basic(fv_tenant_name, fv_bd_name string) string {
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

	`, fv_tenant_name, fv_bd_name)
}

func testAccCheckAciBridgeDomainExists(name string, bridge_domain *models.BridgeDomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Bridge Domain %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Bridge Domain dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		bridge_domainFound := models.BridgeDomainFromContainer(cont)
		if bridge_domainFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Bridge Domain %s not found", rs.Primary.ID)
		}
		*bridge_domain = *bridge_domainFound
		return nil
	}
}

func testAccCheckAciBridgeDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_bridge_domain" {
			cont, err := client.Get(rs.Primary.ID)
			bridge_domain := models.BridgeDomainFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Bridge Domain %s Still exists", bridge_domain.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciBridgeDomainAttributes(fv_tenant_name, fv_bd_name, description string, bridge_domain *models.BridgeDomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_bd_name != GetMOName(bridge_domain.DistinguishedName) {
			return fmt.Errorf("Bad fv_bd %s", GetMOName(bridge_domain.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(bridge_domain.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(bridge_domain.DistinguishedName)))
		}
		if description != bridge_domain.Description {
			return fmt.Errorf("Bad bridge_domain Description %s", bridge_domain.Description)
		}

		return nil
	}
}
