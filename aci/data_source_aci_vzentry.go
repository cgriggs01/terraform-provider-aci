package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciFilterEntry() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciFilterEntryRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"filter_dn": &schema.Schema{
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

func dataSourceAciFilterEntryRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("e-%s", name)
	FilterDn := d.Get("filter_dn").(string)

	dn := fmt.Sprintf("%s/%s", FilterDn, rn)

	vzEntry, err := getRemoteFilterEntry(aciClient, dn)

	if err != nil {
		return err
	}
	setFilterEntryAttributes(vzEntry, d)
	return nil
}
