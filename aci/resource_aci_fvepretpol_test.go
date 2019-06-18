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

func TestAccAciEndPointRetentionPolicy_Basic(t *testing.T) {
	var end_point_retention_policy models.EndPointRetentionPolicy
	fv_tenant_name := acctest.RandString(5)
	fv_ep_ret_pol_name := acctest.RandString(5)
	description := "end_point_retention_policy created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciEndPointRetentionPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciEndPointRetentionPolicyConfig_basic(fv_tenant_name, fv_ep_ret_pol_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciEndPointRetentionPolicyExists("aci_end_point_retention_policy.fooend_point_retention_policy", &end_point_retention_policy),
					testAccCheckAciEndPointRetentionPolicyAttributes(fv_tenant_name, fv_ep_ret_pol_name, description, &end_point_retention_policy),
				),
			},
		},
	})
}

func testAccCheckAciEndPointRetentionPolicyConfig_basic(fv_tenant_name, fv_ep_ret_pol_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_end_point_retention_policy" "fooend_point_retention_policy" {
		name 		= "%s"
		description = "end_point_retention_policy created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, fv_ep_ret_pol_name)
}

func testAccCheckAciEndPointRetentionPolicyExists(name string, end_point_retention_policy *models.EndPointRetentionPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("End Point Retention Policy %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No End Point Retention Policy dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		end_point_retention_policyFound := models.EndPointRetentionPolicyFromContainer(cont)
		if end_point_retention_policyFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("End Point Retention Policy %s not found", rs.Primary.ID)
		}
		*end_point_retention_policy = *end_point_retention_policyFound
		return nil
	}
}

func testAccCheckAciEndPointRetentionPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_end_point_retention_policy" {
			cont, err := client.Get(rs.Primary.ID)
			end_point_retention_policy := models.EndPointRetentionPolicyFromContainer(cont)
			if err == nil {
				return fmt.Errorf("End Point Retention Policy %s Still exists", end_point_retention_policy.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciEndPointRetentionPolicyAttributes(fv_tenant_name, fv_ep_ret_pol_name, description string, end_point_retention_policy *models.EndPointRetentionPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ep_ret_pol_name != GetMOName(end_point_retention_policy.DistinguishedName) {
			return fmt.Errorf("Bad fv_ep_ret_pol %s", GetMOName(end_point_retention_policy.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(end_point_retention_policy.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(end_point_retention_policy.DistinguishedName)))
		}
		if description != end_point_retention_policy.Description {
			return fmt.Errorf("Bad end_point_retention_policy Description %s", end_point_retention_policy.Description)
		}

		return nil
	}
}
