package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/roles"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackRoleAssignment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_role_assignment",
		Description: "Table of all role assignments.",
		List: &plugin.ListConfig{
			Hydrate: listRoleAssignment,
		},
		Columns: []*plugin.Column{
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID").NullIfZero(), Description: "User ID the role is assigned to."},
			{Name: "group_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.ID").NullIfZero(), Description: "Group ID the role is assigned to."},
			{Name: "scope_project_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Scope.Project.ID").NullIfZero(), Description: "Project ID the group/user with role is assigned to."},
			{Name: "scope_domain_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Scope.Domain.ID").NullIfZero(), Description: "Domain ID the group/user with role is assigned to."},
			{Name: "scope_role_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Role.ID").NullIfZero(), Description: "ID of the assigned role."},
		},
	}
}

func listRoleAssignment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_role_assignment.listRoleAssignment", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logger.Error("openstack_role_assignment.listRoleAssignment", "connection_error", err)
		return nil, err
	}

	// get roles
	allPages, err := roles.ListAssignments(identityClient, roles.ListAssignmentsOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_role_assignment.listRoleAssignment", "query_error", err)
		return nil, err
	}

	allRoles, err := roles.ExtractRoleAssignments(allPages)

	for _, role := range allRoles {
		d.StreamListItem(ctx, role)
	}

	return nil, nil
}
