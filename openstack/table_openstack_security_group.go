package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackSecurityGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_security_group",
		Description: "Table of all security groups.",
		List: &plugin.ListConfig{
			Hydrate: listSecurityGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSecurityGroup,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The UUID for the security group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable name for the security group. Might not be unique."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The security group description."},
			{Name: "rules", Type: proto.ColumnType_STRING, Description: "A slice of security group rules that dictate the permitted behaviour for traffic entering and leaving the group."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the security group."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "UpdatedAt contains an ISO-8601 timestamp of when the state of the security group last changed."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "CreatedAt contains an ISO-8601 timestamp of when the security group was created."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the security group."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags optionally set via extensions/attributestags."}},
	}
}

func listSecurityGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_security_group.listSecurityGroup", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_security_group.listSecurityGroup", "connection_error", err)
		return nil, err
	}

	// get security groups
	allPages, err := groups.List(networkClient, groups.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_security_group.listSecurityGroup", "query_error", err)
		return nil, err
	}

	allGroups, err := groups.ExtractGroups(allPages)
	for _, group := range allGroups {
		d.StreamListItem(ctx, group)
	}

	return nil, nil
}

func getSecurityGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_security_group.getSecurityGroup", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get security group
	group, err := groups.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_security_group.getSecurityGroup", "query_error", err)
		return nil, err
	}

	return group, nil
}
