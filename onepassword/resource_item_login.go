package onepassword

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceItemLogin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemLoginRead,
		Create: resourceItemLoginCreate,
		Delete: resourceItemDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				err := resourceItemLoginRead(d, meta)
				return []*schema.ResourceData{d}, err
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vault": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"section": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     sectionSchema(),
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: urlValidate,
			},
		},
	}
}

func resourceItemLoginRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	vaultID := d.Get("vault").(string)
	v, err := m.onePassClient.ReadItem(getID(d), vaultID)
	if err != nil {
		return err
	}
	if v.Template != Category2Template(LoginCategory) {
		return errors.New("item is not from " + string(LoginCategory))
	}

	d.SetId(v.UUID)
	if err := d.Set("name", v.Overview.Title); err != nil {
		return err
	}
	if err := d.Set("url", v.Overview.URL); err != nil {
		return err
	}
	if err := d.Set("notes", v.Details.Notes); err != nil {
		return err
	}
	if err := d.Set("tags", v.Overview.Tags); err != nil {
		return err
	}
	if err := d.Set("vault", v.Vault); err != nil {
		return err
	}
	for _, field := range v.Details.Fields {
		if field.Name == "username" {
			if err := d.Set("username", field.Value); err != nil {
				return err
			}
		}
		if field.Name == "password" {
			if err := d.Set("password", field.Value); err != nil {
				return err
			}
		}
	}
	return d.Set("section", ProcessSections(v.Details.Sections))
}

func resourceItemLoginCreate(d *schema.ResourceData, meta interface{}) error {
	item := &Item{
		Vault:    d.Get("vault").(string),
		Template: Category2Template(LoginCategory),
		Overview: Overview{
			Title: d.Get("name").(string),
			URL:   d.Get("url").(string),
			Tags:  ParseTags(d),
		},
		Details: Details{
			Notes: d.Get("notes").(string),
			Fields: []Field{
				{
					Name:        "username",
					Designation: "username",
					Value:       d.Get("username").(string),
					Type:        FieldText,
				},
				{
					Name:        "password",
					Designation: "password",
					Value:       d.Get("password").(string),
					Type:        FieldPassword,
				},
			},
			Sections: ParseSections(d),
		},
	}
	m := meta.(*Meta)
	err := m.onePassClient.CreateItem(item)
	if err != nil {
		return err
	}
	d.SetId(item.UUID)
	return nil
}
