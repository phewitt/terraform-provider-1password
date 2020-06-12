package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemIdentity() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemIdentityRead,
		Schema: resourceItemIdentity().Schema,
	}
}
