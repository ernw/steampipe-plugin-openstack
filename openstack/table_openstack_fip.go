package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackFip(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_fip",
		Description: "Table of all floating IPs.",
		List: &plugin.ListConfig{
			Hydrate: listFip,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFip,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique identifier for the floating IP instance."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the floating IP instance."},
			{Name: "floating_network_id", Type: proto.ColumnType_STRING, Description: "FloatingNetworkID is the UUID of the external network where the floating IP is to be created."},
			{Name: "floating_ip", Type: proto.ColumnType_STRING, Description: "FloatingIP is the address of the floating IP on the external network."},
			{Name: "port_id", Type: proto.ColumnType_STRING, Description: "PortID is the UUID of the port on an internal network that is associated with the floating IP."},
			{Name: "fixed_ip", Type: proto.ColumnType_STRING, Description: "FixedIP is the specific IP address of the internal port which should be associated with the floating IP."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the floating IP. Only admin users can specify a project identifier other than its own."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "UpdatedAt contains an ISO-8601 timestamp of when the state of the floating ip last changed."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "CreatedAt contains an ISO-8601 timestamp of when the floating ip was created."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the floating IP."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status is the condition of the API resource."},
			{Name: "router_id", Type: proto.ColumnType_STRING, Description: "RouterID is the ID of the router used for this floating IP."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags optionally set via extensions/attributestags."},
		},
	}
}

func listFip(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_fip.listFip", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_fip.listFip", "connection_error", err)
		return nil, err
	}

	// get floating IPs (fips)
	allPages, err := floatingips.List(networkClient, floatingips.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_fip.listFip", "query_error", err)
		return nil, err
	}

	allFIPs, err := floatingips.ExtractFloatingIPs(allPages)
	for _, fip := range allFIPs {
		d.StreamListItem(ctx, fip)
	}

	return nil, nil
}

func getFip(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_fip.getFip", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get floating IP (fip)
	fip, err := floatingips.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_fip.getFip", "query_error", err)
		return nil, err
	}

	return fip, nil
}
