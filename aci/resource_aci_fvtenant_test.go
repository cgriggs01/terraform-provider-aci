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

func TestAccAciTenant_Basic(t *testing.T) {
	var tenant models.Tenant
	fv_tenant_name := acctest.RandString(5)
	description := "tenant created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciTenantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciTenantConfig_basic(fv_tenant_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciTenantExists("aci_tenant.footenant", &tenant),
					testAccCheckAciTenantAttributes(fv_tenant_name, description, &tenant),
				),
			},
		},
	})
}

func testAccCheckAciTenantConfig_basic(fv_tenant_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	`, fv_tenant_name)
}

func testAccCheckAciTenantExists(name string, tenant *models.Tenant) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Tenant %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Tenant dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		tenantFound := models.TenantFromContainer(cont)
		if tenantFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Tenant %s not found", rs.Primary.ID)
		}
		*tenant = *tenantFound
		return nil
	}
}

func testAccCheckAciTenantDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_tenant" {
			cont, err := client.Get(rs.Primary.ID)
			tenant := models.TenantFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Tenant %s Still exists", tenant.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciTenantAttributes(fv_tenant_name, description string, tenant *models.Tenant) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_tenant_name != GetMOName(tenant.DistinguishedName) {
			return fmt.Errorf("Bad fv_tenant %s", GetMOName(tenant.DistinguishedName))
		}

		if description != tenant.Description {
			return fmt.Errorf("Bad tenant Description %s", tenant.Description)
		}

		return nil
	}
}
