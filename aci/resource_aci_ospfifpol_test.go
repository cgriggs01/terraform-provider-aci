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

func TestAccAciOSPFInterfacePolicy_Basic(t *testing.T) {
	var ospf_interface_policy models.OSPFInterfacePolicy
	fv_tenant_name := acctest.RandString(5)
	ospf_if_pol_name := acctest.RandString(5)
	description := "ospf_interface_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciOSPFInterfacePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciOSPFInterfacePolicyConfig_basic(fv_tenant_name, ospf_if_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciOSPFInterfacePolicyExists("aci_ospf_interface_policy.fooospf_interface_policy", &ospf_interface_policy),
					testAccCheckAciOSPFInterfacePolicyAttributes(fv_tenant_name, ospf_if_pol_name, description, &ospf_interface_policy),
				),
			},
		},
	})
}

func testAccCheckAciOSPFInterfacePolicyConfig_basic(fv_tenant_name, ospf_if_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_ospf_interface_policy" "fooospf_interface_policy" {
		name 		= "%s"
		description = "ospf_interface_policy created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, ospf_if_pol_name)
}

func testAccCheckAciOSPFInterfacePolicyExists(name string, ospf_interface_policy *models.OSPFInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("OSPF Interface Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OSPF Interface Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		ospf_interface_policyFound := models.OSPFInterfacePolicyFromContainer(cont)
		if ospf_interface_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("OSPF Interface Policy %s not found", rs.Primary.ID)
		}
		*ospf_interface_policy = *ospf_interface_policyFound
		return nil
	}
}

func testAccCheckAciOSPFInterfacePolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_ospf_interface_policy" {
			cont, err := client.Get(rs.Primary.ID)
			ospf_interface_policy := models.OSPFInterfacePolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("OSPF Interface Policy %s Still exists", ospf_interface_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciOSPFInterfacePolicyAttributes(fv_tenant_name, ospf_if_pol_name, description string, ospf_interface_policy *models.OSPFInterfacePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if ospf_if_pol_name != GetMOName(ospf_interface_policy.DistinguishedName) {
			return fmt.Errorf("Bad ospf_if_pol %s", GetMOName(ospf_interface_policy.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(ospf_interface_policy.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(ospf_interface_policy.DistinguishedName)))
		}
		if description != ospf_interface_policy.Description {
			return fmt.Errorf("Bad ospf_interface_policy Description %s", ospf_interface_policy.Description)
		}

		return nil
	}
}
