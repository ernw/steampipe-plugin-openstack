package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackRouter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_router",
		Description: "Table of all routers.",
		List: &plugin.ListConfig{
			Hydrate: listRouter,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getRouter,
		},
		Columns: []*plugin.Column{
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status indicates whether or not a router is currently operational."},
			{Name: "gateway_info_network_ID", Type: proto.ColumnType_STRING, Transform: transform.FromField("GatewayInfo.NetworkID"), Description: "GatewayInfo represents the information of an external gateway for any particular network router."},
			{Name: "gateway_info_enable_SNAT", Type: proto.ColumnType_BOOL, Transform: transform.FromField("GatewayInfo.EnableSNAT"), Description: "GatewayInfo represents the information of an external gateway for any particular network router."},
			// next item potentially not working correctly here
			{Name: "gateway_info_external_fixed_IPs", Type: proto.ColumnType_JSON, Transform: transform.FromField("GatewayInfo.ExternalFixedIP"), Description: "GatewayInfo represents the information of an external gateway for any particular network router."},
			{Name: "admin_state_up", Type: proto.ColumnType_STRING, Description: "AdminStateUp is the administrative state of the router."},
			{Name: "distributed", Type: proto.ColumnType_BOOL, Description: "Distributed is whether router is disitrubted or not."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the human readable name for the router. It does not have to be unique."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the router."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique identifier for the router."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the router. Only admin users can specify a project identifier other than its own."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the router."},
			{Name: "routes", Type: proto.ColumnType_JSON, Description: "Routes are a collection of static routes that the router will host."},
			{Name: "availability_zone_hints", Type: proto.ColumnType_JSON, Description: "Availability zone hints groups network nodes that run services like DHCP, L3, FW, and others. Used to make network resources highly available."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags optionally set via extensions/attributestags."},
		},
	}
}

func listRouter(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_router.listRouter", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_router.listRouter", "connection_error", err)
		return nil, err
	}

	// get routers
	allPages, err := routers.List(networkClient, routers.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_router.listRouter", "query_error", err)
		return nil, err
	}

	allRouters, err := routers.ExtractRouters(allPages)
	for _, router := range allRouters {
		d.StreamListItem(ctx, router)
	}

	return nil, nil
}

func getRouter(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_router.getRouter", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get router
	router, err := routers.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_router.getRouter", "query_error", err)
		return nil, err
	}

	return router, nil
}
