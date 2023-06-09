package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackServerGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_server_group",
		Description: "Table of all server groups.",
		List: &plugin.ListConfig{
			Hydrate: listServerGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getServerGroup,
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

func listServerGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_server_group.listServerGroup", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_server_group.listServerGroup", "connection_error", err)
		return nil, err
	}

	// get server groups
	allPages, err := servergroups.List(computeClient, servergroups.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_server_group.listServerGroup", "query_error", err)
		return nil, err
	}

	allServergroups, err := servergroups.ExtractServerGroups(allPages)
	for _, servergroup := range allServergroups {
		d.StreamListItem(ctx, servergroup)
	}

	return nil, nil
}

func getServerGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_server_group.getServerGroup", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	// get server group
	servergroup, err := servergroups.Get(computeClient, id).Extract()
	if err != nil {
		logger.Error("openstack_server_group.getServerGroup", "query_error", err)
		return nil, err
	}

	return servergroup, nil
}
