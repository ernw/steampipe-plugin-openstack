package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableOpenstackServergroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_servergroup",
		Description: "Table of all server groups.",
		List: &plugin.ListConfig{
			Hydrate: listServergroup,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of the Server Group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the common name of the server group."},
			{Name: "policies", Type: proto.ColumnType_JSON, Description: "Polices are the group policies."},
			{Name: "members", Type: proto.ColumnType_JSON, Description: "Members are the members of the server group."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "UserID of the server group."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID of the server group."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata includes a list of all user-specified key-value pairs attached to the Server Group."},
			{Name: "policy", Type: proto.ColumnType_STRING, Description: "Policy is the policy of a server group. This requires microversion 2.64 or later."},
			{Name: "max_server_per_host", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rules.MaxServerPerHost"), Description: "MaxServerPerHost specifies how many servers can reside on a single compute host. It can be used only with the \"anti-affinity\" policy."},
		},
	}
}

func listServergroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_servergroup.listServergroup", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_servergroup.listServergroup", "connection_error", err)
		return nil, err
	}

	// get server groups
	allPages, err := servergroups.List(computeClient, servergroups.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_servergroup.listServergroup", "query_error", err)
		return nil, err
	}

	allServergroups, err := servergroups.ExtractServerGroups(allPages)
	for _, servergroup := range allServergroups {
		d.StreamListItem(ctx, servergroup)
	}

	return nil, nil
}
