package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciPortSecurityPolicy() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciPortSecurityPolicyRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciPortSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("infra/portsecurityP-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	l2PortSecurityPol, err := getRemotePortSecurityPolicy(aciClient, dn)

	if err != nil {
		return err
	}
	setPortSecurityPolicyAttributes(l2PortSecurityPol, d)
	return nil
}
