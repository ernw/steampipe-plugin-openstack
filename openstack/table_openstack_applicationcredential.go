package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableOpenstackApplicationcredential(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_applicationcredential",
		Description: "Table of all application credentials.",
		List: &plugin.ListConfig{
			Hydrate: listApplicationcredential,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the application credential."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the application credential."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the application credential’s purpose."},
			{Name: "unrestricted", Type: proto.ColumnType_BOOL, Description: "A flag indicating whether the application credential may be used for creation or destruction of other application credentials or trusts."},
			{Name: "secret", Type: proto.ColumnType_STRING, Description: "Description is the description of the Domain."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project the application credential was created for and that authentication requests using this application credential will be scoped to."},
			{Name: "roles", Type: proto.ColumnType_JSON, Description: "A list of one or more roles that this application credential has associated with its project. A token using this application credential will have these same roles."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "The expiration time of the application credential, if one was specified."},
			{Name: "access_rules", Type: proto.ColumnType_JSON, Description: "A list of access rules objects."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links contains referencing links to the application credential."},
		},
	}
}

func listApplicationcredential(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_applicationcredential.listApplicationcredential", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_applicationcredential.listApplicationcredential", "connection_error", err)
		return nil, err
	}

	// get users
	allPages, err := users.List(identityClient, users.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_applicationcredential.listApplicationcredential", "query_error", err)
		return nil, err
	}

	allUsers, err := users.ExtractUsers(allPages)

	for _, user := range allUsers {

		// get application credentials for current user
		allPages, err := applicationcredentials.List(identityClient, user.ID, applicationcredentials.ListOpts{}).AllPages()
		if err != nil {
			logger.Error("openstack_applicationcredential.listApplicationcredential", "query_error", err)
			return nil, err
		}

		allApplicationcredentials, err := applicationcredentials.ExtractApplicationCredentials(allPages)
		for _, applicationcredential := range allApplicationcredentials {
			d.StreamListItem(ctx, applicationcredential)
		}
	}

	return nil, nil
}
