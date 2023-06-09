package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackApplicationCredential(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_application_credential",
		Description: "Table of all application credentials.",
		List: &plugin.ListConfig{
			Hydrate: listApplicationCredential,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"user_id", "id"}),
			Hydrate:    getApplicationCredential,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ApplicationCredential.ID"), Description: "The ID of the application credential."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "The ID of the user the application credential belongs to."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("ApplicationCredential.Name"), Description: "The name of the application credential."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("ApplicationCredential.Description"), Description: "A description of the application credential’s purpose."},
			{Name: "unrestricted", Type: proto.ColumnType_BOOL, Transform: transform.FromField("ApplicationCredential.Unrestricted"), Description: "A flag indicating whether the application credential may be used for creation or destruction of other application credentials or trusts."},
			{Name: "secret", Type: proto.ColumnType_STRING, Transform: transform.FromField("ApplicationCredential.Secret"), Description: "Description is the description of the Domain."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ApplicationCredential.ProjectID"), Description: "The ID of the project the application credential was created for and that authentication requests using this application credential will be scoped to."},
			{Name: "roles", Type: proto.ColumnType_JSON, Transform: transform.FromField("ApplicationCredential.Roles"), Description: "A list of one or more roles that this application credential has associated with its project. A token using this application credential will have these same roles."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ApplicationCredential.ExpiresAt"), Description: "The expiration time of the application credential, if one was specified."},
			{Name: "access_rules", Type: proto.ColumnType_JSON, Transform: transform.FromField("ApplicationCredential.AccessRules"), Description: "A list of access rules objects."},
			{Name: "links", Type: proto.ColumnType_JSON, Transform: transform.FromField("ApplicationCredential.Links"), Description: "Links contains referencing links to the application credential."},
		},
	}
}

type ApplicationCredentialEntry struct {
	ApplicationCredential applicationcredentials.ApplicationCredential
	UserID                string
}

func listApplicationCredential(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_application_credential.listApplicationCredential", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_application_credential.listApplicationCredential", "connection_error", err)
		return nil, err
	}

	// get users
	allPages, err := users.List(identityClient, users.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_application_credential.listApplicationCredential", "query_error", err)
		return nil, err
	}

	allUsers, err := users.ExtractUsers(allPages)

	for _, user := range allUsers {
		// get application credentials for current user
		allPages, err := applicationcredentials.List(identityClient, user.ID, applicationcredentials.ListOpts{}).AllPages()
		if err != nil {
			logger.Error("openstack_application_credential.listApplicationCredential", "query_error", err)
			return nil, err
		}

		allApplicationcredentials, err := applicationcredentials.ExtractApplicationCredentials(allPages)
		for _, applicationcredential := range allApplicationcredentials {
			d.StreamListItem(ctx, ApplicationCredentialEntry{applicationcredential, user.ID})
		}
	}

	return nil, nil
}

func getApplicationCredential(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()
	uid := d.EqualsQuals["user_id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_application_credential.getApplicationCredential", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	// get application credential for provided user id (user_id) and application credential id
	applicationcredential, err := applicationcredentials.Get(identityClient, uid, id).Extract()
	if err != nil {
		logger.Error("openstack_application_credential.getApplicationCredential", "query_error", err)
		return nil, err
	}

	return ApplicationCredentialEntry{*applicationcredential, uid}, nil
}
