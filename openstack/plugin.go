package openstack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-openstack",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"openstack_aggregate":              tableOpenstackAggregate(ctx),
			"openstack_application_credential": tableOpenstackApplicationCredential(ctx),
			"openstack_availability_zone":      tableOpenstackAvailabilityZone(ctx),
			"openstack_compute_image":          tableOpenstackComputeImage(ctx),
			"openstack_domain":                 tableOpenstackDomain(ctx),
			"openstack_fip":                    tableOpenstackFip(ctx),
			"openstack_group":                  tableOpenstackGroup(ctx),
			"openstack_keypair":                tableOpenstackKeypair(ctx),
			"openstack_network":                tableOpenstackNetwork(ctx),
			"openstack_port":                   tableOpenstackPort(ctx),
			"openstack_project":                tableOpenstackProject(ctx),
			"openstack_role_assignment":        tableOpenstackRoleAssignment(ctx),
			"openstack_role":                   tableOpenstackRole(ctx),
			"openstack_router":                 tableOpenstackRouter(ctx),
			"openstack_security_group_rule":    tableOpenstackSecurityGroupRule(ctx),
			"openstack_security_group":         tableOpenstackSecurityGroup(ctx),
			"openstack_server_group":           tableOpenstackServerGroup(ctx),
			"openstack_server":                 tableOpenstackServer(ctx),
			"openstack_snapshot":               tableOpenstackSnapshot(ctx),
			"openstack_subnet":                 tableOpenstackSubnet(ctx),
			"openstack_user":                   tableOpenstackUser(ctx),
			"openstack_volume_type":            tableOpenstackVolumeType(ctx),
			"openstack_volume":                 tableOpenstackVolume(ctx),
		},
	}

	return p
}
