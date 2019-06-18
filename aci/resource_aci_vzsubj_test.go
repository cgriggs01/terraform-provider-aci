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

func TestAccAciContractSubject_Basic(t *testing.T) {
	var contract_subject models.ContractSubject
	fv_tenant_name := acctest.RandString(5)
	vz_br_cp_name := acctest.RandString(5)
	vz_subj_name := acctest.RandString(5)
	description := "contract_subject created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciContractSubjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciContractSubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractSubjectExists("aci_contract_subject.foocontract_subject", &contract_subject),
					testAccCheckAciContractSubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description, &contract_subject),
				),
			},
		},
	})
}

func testAccCheckAciContractSubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name string) string {
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

	resource "aci_contract_subject" "foocontract_subject" {
		name 		= "%s"
		description = "contract_subject created while acceptance testing"
		contract_dn = "${aci_contract.foocontract.id}"
	}

	`, fv_tenant_name, vz_br_cp_name, vz_subj_name)
}

func testAccCheckAciContractSubjectExists(name string, contract_subject *models.ContractSubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Contract Subject %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Contract Subject dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		contract_subjectFound := models.ContractSubjectFromContainer(cont)
		if contract_subjectFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Contract Subject %s not found", rs.Primary.ID)
		}
		*contract_subject = *contract_subjectFound
		return nil
	}
}

func testAccCheckAciContractSubjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_contract_subject" {
			cont, err := client.Get(rs.Primary.ID)
			contract_subject := models.ContractSubjectFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Contract Subject %s Still exists", contract_subject.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciContractSubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description string, contract_subject *models.ContractSubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_subj_name != GetMOName(contract_subject.DistinguishedName) {
			return fmt.Errorf("Bad vz_subj %s", GetMOName(contract_subject.DistinguishedName))
		}

		if vz_br_cp_name != GetMOName(GetParentDn(contract_subject.DistinguishedName)) {
			return fmt.Errorf(" Bad vz_br_cp %s", GetMOName(GetParentDn(contract_subject.DistinguishedName)))
		}
		if description != contract_subject.Description {
			return fmt.Errorf("Bad contract_subject Description %s", contract_subject.Description)
		}

		return nil
	}
}
