package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackAvailabilityZone(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_availability_zone",
		Description: "Table of all availability zones.",
		List: &plugin.ListConfig{
			Hydrate: listAvailabilityZone,
		},
		Columns: []*plugin.Column{
			{Name: "hosts", Type: proto.ColumnType_JSON, Description: "Hosts in the availability zone."},
			{Name: "zone_name", Type: proto.ColumnType_STRING, Description: "Name of the availability zone."},
			{Name: "zone_state", Type: proto.ColumnType_JSON, Description: "State of the availability zone."},
		},
	}
}

func listAvailabilityZone(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_availability_zone.listAvailabilityZone", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_availability_zone.listAvailabilityZone", "connection_error", err)
		return nil, err
	}

	// get availability zones
	allPages, err := availabilityzones.List(computeClient).AllPages()
	if err != nil {
		logger.Error("openstack_availability_zone.listAvailabilityZone", "query_error", err)
		return nil, err
	}

	allAvailabilityZones, err := availabilityzones.ExtractAvailabilityZones(allPages)
	for _, availabilityzone := range allAvailabilityZones {
		d.StreamListItem(ctx, availabilityzone)
	}

	return nil, nil
}
