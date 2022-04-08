package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/integrations/terraform-provider-github/v4/github"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: github.Provider})
}
