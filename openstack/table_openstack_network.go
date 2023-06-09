package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackNetwork(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_network",
		Description: "Table of all networks.",
		List: &plugin.ListConfig{
			Hydrate: listNetwork,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getNetwork,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "UUID for the network."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable name for the network. Might not be unique."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the network."},
			{Name: "admin_state_up", Type: proto.ColumnType_BOOL, Description: "The administrative state of network. If false (down), the network does not forward packets."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Indicates whether network is currently operational. Possible values include `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional values."},
			{Name: "subnets", Type: proto.ColumnType_JSON, Description: "Subnets associated with this network."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the network."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "UpdatedAt contains an ISO-8601 timestamp of when the state of the network last changed."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "CreatedAt contains an ISO-8601 timestamp of when the network was created."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the network."},
			{Name: "shared", Type: proto.ColumnType_BOOL, Description: "Specifies whether the network resource can be accessed by any tenant."},
			{Name: "availability_zone_hints", Type: proto.ColumnType_STRING, Description: "Availability zone hints groups network nodes that run services like DHCP, L3, FW, and others. Used to make network resources highly available."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags optionally set via extensions/attributestags."},
			{Name: "revision_number", Type: proto.ColumnType_INT, Description: "RevisionNumber optionally set via extensions/standard-attr-revisions."},
		},
	}
}

func listNetwork(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_network.listNetwork", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_network.listNetwork", "connection_error", err)
		return nil, err
	}

	// get networks
	allPages, err := networks.List(networkClient, networks.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_network.listNetwork", "query_error", err)
		return nil, err
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	for _, network := range allNetworks {
		d.StreamListItem(ctx, network)
	}

	return nil, nil
}

func getNetwork(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_network.getNetwork", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get network
	network, err := networks.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_network.getNetwork", "query_error", err)
		return nil, err
	}

	return network, nil
}
