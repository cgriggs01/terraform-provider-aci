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

func TestAccAciFilterEntry_Basic(t *testing.T) {
	var filter_entry models.FilterEntry
	fv_tenant_name := acctest.RandString(5)
	vz_filter_name := acctest.RandString(5)
	vz_entry_name := acctest.RandString(5)
	description := "filter_entry created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciFilterEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFilterEntryConfig_basic(fv_tenant_name, vz_filter_name, vz_entry_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFilterEntryExists("aci_filter_entry.foofilter_entry", &filter_entry),
					testAccCheckAciFilterEntryAttributes(fv_tenant_name, vz_filter_name, vz_entry_name, description, &filter_entry),
				),
			},
		},
	})
}

func testAccCheckAciFilterEntryConfig_basic(fv_tenant_name, vz_filter_name, vz_entry_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_filter" "foofilter" {
		name 		= "%s"
		description = "filter created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_filter_entry" "foofilter_entry" {
		name 		= "%s"
		description = "filter_entry created while acceptance testing"
		filter_dn = "${aci_filter.foofilter.id}"
	}

	`, fv_tenant_name, vz_filter_name, vz_entry_name)
}

func testAccCheckAciFilterEntryExists(name string, filter_entry *models.FilterEntry) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Filter Entry %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Filter Entry dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		filter_entryFound := models.FilterEntryFromContainer(cont)
		if filter_entryFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Filter Entry %s not found", rs.Primary.ID)
		}
		*filter_entry = *filter_entryFound
		return nil
	}
}

func testAccCheckAciFilterEntryDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_filter_entry" {
			cont, err := client.Get(rs.Primary.ID)
			filter_entry := models.FilterEntryFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Filter Entry %s Still exists", filter_entry.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciFilterEntryAttributes(fv_tenant_name, vz_filter_name, vz_entry_name, description string, filter_entry *models.FilterEntry) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_entry_name != GetMOName(filter_entry.DistinguishedName) {
			return fmt.Errorf("Bad vz_entry %s", GetMOName(filter_entry.DistinguishedName))
		}

		if vz_filter_name != GetMOName(GetParentDn(filter_entry.DistinguishedName)) {
			return fmt.Errorf(" Bad vz_filter %s", GetMOName(GetParentDn(filter_entry.DistinguishedName)))
		}
		if description != filter_entry.Description {
			return fmt.Errorf("Bad filter_entry Description %s", filter_entry.Description)
		}

		return nil
	}
}
