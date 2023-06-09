package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/aggregates"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackAggregate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_aggregate",
		Description: "Table of all host aggregates.",
		List: &plugin.ListConfig{
			Hydrate: listAggregate,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAggregate,
		},
		Columns: []*plugin.Column{
			{Name: "availability_zone", Type: proto.ColumnType_STRING, Description: "The availability zone of the host aggregate."},
			{Name: "hosts", Type: proto.ColumnType_JSON, Description: "A list of host ids in this aggregate."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "The ID of the host aggregate."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata key and value pairs associate with the aggregate."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the aggregate."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the resource was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the resource was updated."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time when the resource was deleted."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "Deleted indicates whether this aggregate is deleted or not."},
		},
	}
}

func listAggregate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_aggregate.listAggregate", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_aggregate.listAggregate", "connection_error", err)
		return nil, err
	}

	// get host aggregates
	allPages, err := aggregates.List(computeClient).AllPages()
	if err != nil {
		logger.Error("openstack_aggregate.listAggregate", "query_error", err)
		return nil, err
	}

	allAggregates, err := aggregates.ExtractAggregates(allPages)
	for _, aggregate := range allAggregates {
		d.StreamListItem(ctx, aggregate)
	}

	return nil, nil
}

func getAggregate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetInt64Value()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_aggregate.getAggregate", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	// get aggregate
	aggregate, err := aggregates.Get(computeClient, int(id)).Extract()
	if err != nil {
		logger.Error("openstack_aggregate.getAggregate", "query_error", err)
		return nil, err
	}

	return aggregate, nil
}
