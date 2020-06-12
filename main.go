package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/phewitt/terraform-provider-1password/onepassword"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onepassword.Provider,
	})
}
