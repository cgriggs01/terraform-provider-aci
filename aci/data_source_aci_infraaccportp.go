package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciLeafInterfaceProfile() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciLeafInterfaceProfileRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciLeafInterfaceProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/accportprof-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	infraAccPortP, err := getRemoteLeafInterfaceProfile(aciClient, dn)

	if err != nil {
		return err
	}
	setLeafInterfaceProfileAttributes(infraAccPortP, d)
	return nil
}
