package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemDocument() *schema.Resource {
	s := resourceItemDocument().Schema
	delete(s, "file_path")

	return &schema.Resource{
		Read:   resourceItemDocumentRead,
		Schema: s,
	}
}
