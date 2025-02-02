package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciTenant() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciTenantRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciTenantRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("tn-%s", name)

	dn := fmt.Sprintf("uni/%s", rn)

	fvTenant, err := getRemoteTenant(aciClient, dn)

	if err != nil {
		return err
	}
	setTenantAttributes(fvTenant, d)
	return nil
}
