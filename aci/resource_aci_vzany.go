package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciAny() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciAnyCreate,
		Update: resourceAciAnyUpdate,
		Read:   resourceAciAnyRead,
		Delete: resourceAciAnyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciAnyImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"vrf_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"match_t": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pref_gr_memb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"relation_vz_rs_any_to_cons": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			"relation_vz_rs_any_to_cons_if": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			"relation_vz_rs_any_to_prov": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
		}),
	}
}

func getRemoteAny(client *client.Client, dn string) (*models.Any, error) {
	vzAnyCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzAny := models.AnyFromContainer(vzAnyCont)

	if vzAny.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzAny.DistinguishedName)
	}

	return vzAny, nil
}

func setAnyAttributes(vzAny *models.Any, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzAny.DistinguishedName)
	d.Set("description", vzAny.Description)
	d.Set("vrf_dn", GetParentDn(vzAny.DistinguishedName))
	vzAnyMap, _ := vzAny.ToMap()

	d.Set("annotation", vzAnyMap["annotation"])
	d.Set("match_t", vzAnyMap["matchT"])
	d.Set("name_alias", vzAnyMap["nameAlias"])
	d.Set("pref_gr_memb", vzAnyMap["prefGrMemb"])
	return d
}

func resourceAciAnyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzAny, err := getRemoteAny(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setAnyAttributes(vzAny, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciAnyCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	VRFDn := d.Get("vrf_dn").(string)

	vzAnyAttr := models.AnyAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzAnyAttr.Annotation = Annotation.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		vzAnyAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzAnyAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		vzAnyAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	vzAny := models.NewAny(fmt.Sprintf("any"), VRFDn, desc, vzAnyAttr)

	err := aciClient.Save(vzAny)
	if err != nil {
		return err
	}

	if relationTovzRsAnyToCons, ok := d.GetOk("relation_vz_rs_any_to_cons"); ok {
		relationParamList := toStringList(relationTovzRsAnyToCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationvzRsAnyToConsFromAny(vzAny.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTovzRsAnyToConsIf, ok := d.GetOk("relation_vz_rs_any_to_cons_if"); ok {
		relationParamList := toStringList(relationTovzRsAnyToConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationvzRsAnyToConsIfFromAny(vzAny.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTovzRsAnyToProv, ok := d.GetOk("relation_vz_rs_any_to_prov"); ok {
		relationParamList := toStringList(relationTovzRsAnyToProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationvzRsAnyToProvFromAny(vzAny.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(vzAny.DistinguishedName)
	return resourceAciAnyRead(d, m)
}

func resourceAciAnyUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VRFDn := d.Get("vrf_dn").(string)

	vzAnyAttr := models.AnyAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzAnyAttr.Annotation = Annotation.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		vzAnyAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzAnyAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		vzAnyAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	vzAny := models.NewAny(fmt.Sprintf("any"), VRFDn, desc, vzAnyAttr)

	vzAny.Status = "modified"

	err := aciClient.Save(vzAny)

	if err != nil {
		return err
	}

	if d.HasChange("relation_vz_rs_any_to_cons") {
		oldRel, newRel := d.GetChange("relation_vz_rs_any_to_cons")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationvzRsAnyToConsFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationvzRsAnyToConsFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_vz_rs_any_to_cons_if") {
		oldRel, newRel := d.GetChange("relation_vz_rs_any_to_cons_if")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationvzRsAnyToConsIfFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationvzRsAnyToConsIfFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_vz_rs_any_to_prov") {
		oldRel, newRel := d.GetChange("relation_vz_rs_any_to_prov")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationvzRsAnyToProvFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationvzRsAnyToProvFromAny(vzAny.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(vzAny.DistinguishedName)
	return resourceAciAnyRead(d, m)

}

func resourceAciAnyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzAny, err := getRemoteAny(aciClient, dn)

	if err != nil {
		return err
	}
	setAnyAttributes(vzAny, d)
	return nil
}

func resourceAciAnyDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzAny")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
