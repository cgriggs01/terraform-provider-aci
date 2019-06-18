package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciInterfaceFCPolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciInterfaceFCPolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciInterfaceFCPolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/fcIfPol-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	fcIfPol, err := getRemoteInterfaceFCPolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setInterfaceFCPolicyAttributes(fcIfPol, d)
	return nil
}
