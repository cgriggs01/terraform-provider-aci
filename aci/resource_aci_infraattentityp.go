package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciAttachableAccessEntityProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciAttachableAccessEntityProfileCreate,
		Update: resourceAciAttachableAccessEntityProfileUpdate,
		Read:   resourceAciAttachableAccessEntityProfileRead,
		Delete: resourceAciAttachableAccessEntityProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciAttachableAccessEntityProfileImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"relation_infra_rs_dom_p": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
		}),
	}
}
func getRemoteAttachableAccessEntityProfile(client *client.Client, dn string) (*models.AttachableAccessEntityProfile, error) {
	infraAttEntityPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	infraAttEntityP := models.AttachableAccessEntityProfileFromContainer(infraAttEntityPCont)

	if infraAttEntityP.DistinguishedName == "" {
		return nil, fmt.Errorf("AttachableAccessEntityProfile %s not found", infraAttEntityP.DistinguishedName)
	}

	return infraAttEntityP, nil
}

func setAttachableAccessEntityProfileAttributes(infraAttEntityP *models.AttachableAccessEntityProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(infraAttEntityP.DistinguishedName)
	d.Set("description", infraAttEntityP.Description)
	d.Set("name", GetMOName(infraAttEntityP.DistinguishedName))
	infraAttEntityPMap, _ := infraAttEntityP.ToMap()

	d.Set("annotation", infraAttEntityPMap["annotation"])
	d.Set("name_alias", infraAttEntityPMap["nameAlias"])
	return d
}

func resourceAciAttachableAccessEntityProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	infraAttEntityP, err := getRemoteAttachableAccessEntityProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setAttachableAccessEntityProfileAttributes(infraAttEntityP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciAttachableAccessEntityProfileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	infraAttEntityPAttr := models.AttachableAccessEntityProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraAttEntityPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraAttEntityPAttr.NameAlias = NameAlias.(string)
	}
	infraAttEntityP := models.NewAttachableAccessEntityProfile(fmt.Sprintf("infra/attentp-%s", name), "uni", desc, infraAttEntityPAttr)

	err := aciClient.Save(infraAttEntityP)
	if err != nil {
		return err
	}

	if relationToinfraRsDomP, ok := d.GetOk("relation_infra_rs_dom_p"); ok {
		relationParamList := toStringList(relationToinfraRsDomP.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationinfraRsDomPFromAttachableAccessEntityProfile(infraAttEntityP.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(infraAttEntityP.DistinguishedName)
	return resourceAciAttachableAccessEntityProfileRead(d, m)
}

func resourceAciAttachableAccessEntityProfileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	infraAttEntityPAttr := models.AttachableAccessEntityProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraAttEntityPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraAttEntityPAttr.NameAlias = NameAlias.(string)
	}
	infraAttEntityP := models.NewAttachableAccessEntityProfile(fmt.Sprintf("infra/attentp-%s", name), "uni", desc, infraAttEntityPAttr)

	infraAttEntityP.Status = "modified"

	err := aciClient.Save(infraAttEntityP)

	if err != nil {
		return err
	}

	if d.HasChange("relation_infra_rs_dom_p") {
		oldRel, newRel := d.GetChange("relation_infra_rs_dom_p")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationinfraRsDomPFromAttachableAccessEntityProfile(infraAttEntityP.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationinfraRsDomPFromAttachableAccessEntityProfile(infraAttEntityP.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(infraAttEntityP.DistinguishedName)
	return resourceAciAttachableAccessEntityProfileRead(d, m)

}

func resourceAciAttachableAccessEntityProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	infraAttEntityP, err := getRemoteAttachableAccessEntityProfile(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setAttachableAccessEntityProfileAttributes(infraAttEntityP, d)

	infraRsDomPData, err := aciClient.ReadRelationinfraRsDomPFromAttachableAccessEntityProfile(dn)
	if err != nil {
		log.Printf("[DEBUG] Error while reading relation infraRsDomP %v", err)

	} else {
		d.Set("relation_infra_rs_dom_p", infraRsDomPData)
	}

	return nil
}

func resourceAciAttachableAccessEntityProfileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "infraAttEntityP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
