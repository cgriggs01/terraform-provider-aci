package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciApplicationEPG() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciApplicationEPGRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"application_profile_dn": &schema.Schema{
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

func dataSourceAciApplicationEPGRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("epg-%s", name)
	ApplicationProfileDn := d.Get("application_profile_dn").(string)

	dn := fmt.Sprintf("%s/%s", ApplicationProfileDn, rn)

	fvAEPg, err := getRemoteApplicationEPG(aciClient, dn)

	if err != nil {
		return err
	}
	setApplicationEPGAttributes(fvAEPg, d)
	return nil
}
