package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciVMMDomain() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciVMMDomainRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"provider_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciVMMDomainRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("dom-%s", name)
	ProviderProfileDn := d.Get("provider_profile_dn").(string)

	dn := fmt.Sprintf("%s/%s", ProviderProfileDn, rn)

	vmmDomP, err := getRemoteVMMDomain(aciClient, dn)

	if err != nil {
		return err
	}
	setVMMDomainAttributes(vmmDomP, d)
	return nil
}
