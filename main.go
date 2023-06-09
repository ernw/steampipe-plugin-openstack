package main

import (
	"github.com/turbot/steampipe-plugin-openstack/openstack"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: openstack.Plugin})
}
