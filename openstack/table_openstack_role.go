package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_role",
		Description: "Table of all roles.",
		List: &plugin.ListConfig{
			Hydrate: listRole,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRole,
		},
		Columns: []*plugin.Column{
			{Name: "domain_id", Type: proto.ColumnType_STRING, Description: "DomainID is the domain ID the role belongs to."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of the role."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the role name."},
			{Name: "extra", Type: proto.ColumnType_JSON, Description: "Description is the description of the role."},
		},
	}
}

func listRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_role.listRole", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_role.listRole", "connection_error", err)
		return nil, err
	}

	// get users
	allPages, err := roles.List(identityClient, roles.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_role.listRole", "query_error", err)
		return nil, err
	}

	allRoles, err := roles.ExtractRoles(allPages)
	for _, role := range allRoles {
		d.StreamListItem(ctx, role)
	}

	return nil, nil
}

func getRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_role.getRole", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	// get role
	role, err := roles.Get(identityClient, id).Extract()
	if err != nil {
		logger.Error("openstack_role.getRole", "query_error", err)
		return nil, err
	}

	return role, nil
}
