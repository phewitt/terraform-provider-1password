package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemPassword() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemPasswordRead,
		Schema: resourceItemPassword().Schema,
	}
}
