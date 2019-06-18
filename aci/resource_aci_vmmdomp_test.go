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

func TestAccAciVMMDomain_Basic(t *testing.T) {
	var vmm_domain models.VMMDomain
	vmm_prov_p_name := acctest.RandString(5)
	vmm_dom_p_name := acctest.RandString(5)
	description := "vmm_domain created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciVMMDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciVMMDomainConfig_basic(vmm_prov_p_name, vmm_dom_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciVMMDomainExists("aci_vmm_domain.foovmm_domain", &vmm_domain),
					testAccCheckAciVMMDomainAttributes(vmm_prov_p_name, vmm_dom_p_name, description, &vmm_domain),
				),
			},
		},
	})
}

func testAccCheckAciVMMDomainConfig_basic(vmm_prov_p_name, vmm_dom_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_provider_profile" "fooprovider_profile" {
		name 		= "%s"
		description = "provider_profile created while acceptance testing"

	}

	resource "aci_vmm_domain" "foovmm_domain" {
		name 		= "%s"
		description = "vmm_domain created while acceptance testing"
		provider_profile_dn = "${aci_provider_profile.fooprovider_profile.id}"
	}

	`, vmm_prov_p_name, vmm_dom_p_name)
}

func testAccCheckAciVMMDomainExists(name string, vmm_domain *models.VMMDomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VMM Domain %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VMM Domain dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		vmm_domainFound := models.VMMDomainFromContainer(cont)
		if vmm_domainFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("VMM Domain %s not found", rs.Primary.ID)
		}
		*vmm_domain = *vmm_domainFound
		return nil
	}
}

func testAccCheckAciVMMDomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_vmm_domain" {
			cont, err := client.Get(rs.Primary.ID)
			vmm_domain := models.VMMDomainFromContainer(cont)
			if err == nil {
				return fmt.Errorf("VMM Domain %s Still exists", vmm_domain.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciVMMDomainAttributes(vmm_prov_p_name, vmm_dom_p_name, description string, vmm_domain *models.VMMDomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vmm_dom_p_name != GetMOName(vmm_domain.DistinguishedName) {
			return fmt.Errorf("Bad vmm_dom_p %s", GetMOName(vmm_domain.DistinguishedName))
		}

		if vmm_prov_p_name != GetMOName(GetParentDn(vmm_domain.DistinguishedName)) {
			return fmt.Errorf(" Bad vmm_prov_p %s", GetMOName(GetParentDn(vmm_domain.DistinguishedName)))
		}
		if description != vmm_domain.Description {
			return fmt.Errorf("Bad vmm_domain Description %s", vmm_domain.Description)
		}

		return nil
	}
}
