package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_project",
		Description: "Table of all projects.",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the name of the project."},
			{Name: "is_domain", Type: proto.ColumnType_BOOL, Description: "IsDomain indicates whether the project is a domain."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description is the description of the project."},
			{Name: "domain_id", Type: proto.ColumnType_STRING, Description: "DomainID is the domain ID the project belongs to."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "Enabled is whether or not the project is enabled."},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Description: "ParentID is the parent_id of the project."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags is the list of tags associated with the project."},
			{Name: "options", Type: proto.ColumnType_STRING, Description: "Options are defined options in the API to enable certain features."},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_project.listProject", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_project.listProject", "connection_error", err)
		return nil, err
	}

	// get projects
	allPages, err := projects.List(identityClient, projects.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_project.listProject", "query_error", err)
		return nil, err
	}

	allProjects, err := projects.ExtractProjects(allPages)
	for _, project := range allProjects {
		d.StreamListItem(ctx, project)
	}

	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_project.getProject", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	// get project
	project, err := projects.Get(identityClient, id).Extract()
	if err != nil {
		logger.Error("openstack_project.getProject", "query_error", err)
		return nil, err
	}

	return project, nil
}
