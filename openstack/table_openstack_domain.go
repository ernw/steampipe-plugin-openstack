package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/domains"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_domain",
		Description: "Table of all domains.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDomain,
		},
		Columns: []*plugin.Column{
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description is the description of the Domain."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "Enabled is whether or not the domain is enabled."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of the domain."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links contains referencing links to the domain."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the name of the domain."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_domain.listDomain", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_domain.listDomain", "connection_error", err)
		return nil, err
	}

	// get domains
	allPages, err := domains.List(identityClient, domains.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_domain.listDomain", "query_error", err)
		return nil, err
	}

	allDomains, err := domains.ExtractDomains(allPages)
	for _, domain := range allDomains {
		d.StreamListItem(ctx, domain)
	}

	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_domain.getDomain", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	// get domain
	domain, err := domains.Get(identityClient, id).Extract()
	if err != nil {
		logger.Error("openstack_domain.getDomain", "query_error", err)
		return nil, err
	}

	return domain, nil
}
