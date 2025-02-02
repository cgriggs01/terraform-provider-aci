package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciMiscablingProtocolInterfacePolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciMiscablingProtocolInterfacePolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciMiscablingProtocolInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/mcpIfP-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	mcpIfPol, err := getRemoteMiscablingProtocolInterfacePolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setMiscablingProtocolInterfacePolicyAttributes(mcpIfPol, d)
	return nil
}
