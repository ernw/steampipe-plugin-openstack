package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/snapshots"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackSnapshot(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_snapshot",
		Description: "Table of all volume snapshots.",
		List: &plugin.ListConfig{
			Hydrate: listSnapshot,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSnapshot,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date when this snapshot was last updated."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date when this snapshot was created."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable display name for the snapshot."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Human-readable description for the snapshot."},
			{Name: "volume_id", Type: proto.ColumnType_STRING, Description: "ID of the Volume from which this Snapshot was created."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the snapshot."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the snapshot in GB."},
			{Name: "metadata", Type: proto.ColumnType_STRING, Description: "User-defined key-value pairs."},
		},
	}
}

func listSnapshot(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_snapshot.listSnapshot", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_snapshot.listSnapshot", "connection_error", err)
		return nil, err
	}

	// get volume snapshots
	allPages, err := snapshots.List(blockStorageClient, snapshots.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_snapshot.listSnapshot", "query_error", err)
		return nil, err
	}

	allSnapshots, err := snapshots.ExtractSnapshots(allPages)
	for _, snapshot := range allSnapshots {

		// gophercloud requires to extract details for each snapshot
		snapshotDetails, err := snapshots.Get(blockStorageClient, snapshot.ID).Extract()
		if err != nil {
			logger.Error("openstack_snapshot.listSnapshot", "query_error", err)
		}
		d.StreamListItem(ctx, snapshotDetails)
	}

	return nil, nil
}

func getSnapshot(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_snapshot.getSnapshot", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	// get snapshot
	snapshot, err := snapshots.Get(blockStorageClient, id).Extract()
	if err != nil {
		logger.Error("openstack_snapshot.getSnapshot", "query_error", err)
		return nil, err
	}

	return snapshot, nil
}
