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

func TestAccAciContract_Basic(t *testing.T) {
	var contract models.Contract
	fv_tenant_name := acctest.RandString(5)
	vz_br_cp_name := acctest.RandString(5)
	description := "contract created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciContractConfig_basic(fv_tenant_name, vz_br_cp_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists("aci_contract.foocontract", &contract),
					testAccCheckAciContractAttributes(fv_tenant_name, vz_br_cp_name, description, &contract),
				),
			},
		},
	})
}

func testAccCheckAciContractConfig_basic(fv_tenant_name, vz_br_cp_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_contract" "foocontract" {
		name 		= "%s"
		description = "contract created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, vz_br_cp_name)
}

func testAccCheckAciContractExists(name string, contract *models.Contract) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Contract %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Contract dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		contractFound := models.ContractFromContainer(cont)
		if contractFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Contract %s not found", rs.Primary.ID)
		}
		*contract = *contractFound
		return nil
	}
}

func testAccCheckAciContractDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_contract" {
			cont, err := client.Get(rs.Primary.ID)
			contract := models.ContractFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Contract %s Still exists", contract.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciContractAttributes(fv_tenant_name, vz_br_cp_name, description string, contract *models.Contract) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_br_cp_name != GetMOName(contract.DistinguishedName) {
			return fmt.Errorf("Bad vz_br_cp %s", GetMOName(contract.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(contract.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(contract.DistinguishedName)))
		}
		if description != contract.Description {
			return fmt.Errorf("Bad contract Description %s", contract.Description)
		}

		return nil
	}
}
