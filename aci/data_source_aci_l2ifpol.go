package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciL2InterfacePolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciL2InterfacePolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciL2InterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/l2IfP-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	l2IfPol, err := getRemoteL2InterfacePolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setL2InterfacePolicyAttributes(l2IfPol, d)
	return nil
}
