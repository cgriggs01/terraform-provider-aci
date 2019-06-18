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

func TestAccAciPCVPCInterfacePolicyGroup_Basic(t *testing.T) {
	var pcvpc_interface_policy_group models.PCVPCInterfacePolicyGroup
	infra_acc_bndl_grp_name := acctest.RandString(5)
	description := "pc/vpc_interface_policy_group created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciPCVPCInterfacePolicyGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciPCVPCInterfacePolicyGroupConfig_basic(infra_acc_bndl_grp_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciPCVPCInterfacePolicyGroupExists("aci_pc/vpc_interface_policy_group.foopc/vpc_interface_policy_group", &pcvpc_interface_policy_group),
					testAccCheckAciPCVPCInterfacePolicyGroupAttributes(infra_acc_bndl_grp_name, description, &pcvpc_interface_policy_group),
				),
			},
		},
	})
}

func testAccCheckAciPCVPCInterfacePolicyGroupConfig_basic(infra_acc_bndl_grp_name string) string {
	return fmt.Sprintf(`

	resource "aci_pc/vpc_interface_policy_group" "foopc/vpc_interface_policy_group" {
		name 		= "%s"
		description = "pc/vpc_interface_policy_group created while acceptance testing"

	}

	`, infra_acc_bndl_grp_name)
}

func testAccCheckAciPCVPCInterfacePolicyGroupExists(name string, pcvpc_interface_policy_group *models.PCVPCInterfacePolicyGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("PC/VPC Interface Policy Group %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No PC/VPC Interface Policy Group dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		pcvpc_interface_policy_groupFound := models.PCVPCInterfacePolicyGroupFromContainer(cont)
		if pcvpc_interface_policy_groupFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("PC/VPC Interface Policy Group %s not found", rs.Primary.ID)
		}
		*pcvpc_interface_policy_group = *pcvpc_interface_policy_groupFound
		return nil
	}
}

func testAccCheckAciPCVPCInterfacePolicyGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_pc/vpc_interface_policy_group" {
			cont, err := client.Get(rs.Primary.ID)
			pcvpc_interface_policy_group := models.PCVPCInterfacePolicyGroupFromContainer(cont)
			if err == nil {
				return fmt.Errorf("PC/VPC Interface Policy Group %s Still exists", pcvpc_interface_policy_group.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciPCVPCInterfacePolicyGroupAttributes(infra_acc_bndl_grp_name, description string, pcvpc_interface_policy_group *models.PCVPCInterfacePolicyGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if infra_acc_bndl_grp_name != GetMOName(pcvpc_interface_policy_group.DistinguishedName) {
			return fmt.Errorf("Bad infra_acc_bndl_grp %s", GetMOName(pcvpc_interface_policy_group.DistinguishedName))
		}

		if description != pcvpc_interface_policy_group.Description {
			return fmt.Errorf("Bad pcvpc_interface_policy_group Description %s", pcvpc_interface_policy_group.Description)
		}

		return nil
	}
}
