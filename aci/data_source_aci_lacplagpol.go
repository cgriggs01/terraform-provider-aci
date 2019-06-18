package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciLACPPolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciLACPPolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciLACPPolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/lacplagp-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	lacpLagPol, err := getRemoteLACPPolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setLACPPolicyAttributes(lacpLagPol, d)
	return nil
}
