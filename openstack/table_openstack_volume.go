package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_volume",
		Description: "Table of all volumes.",
		List: &plugin.ListConfig{
			Hydrate: listVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolume,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the volume."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the volume."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the volume in GB."},
			{Name: "availability_zone", Type: proto.ColumnType_STRING, Description: "AvailabilityZone is which availability zone the volume is in."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date when this volume was created."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The date when this volume was last updated."},
			{Name: "attached_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attachments[0].AttachedAt"), Description: "AttachedAt is the time the attachment was created."},
			{Name: "attachment_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attachments[0].AttachmentID"), Description: "ID is the Unique identifier for the attachment."},
			{Name: "device", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attachments[0].Device"), Description: "Name of the attached device."},
			{Name: "host_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attachments[0].HostName"), Description: "Attachment Hostname."},
			{Name: "server_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attachments[0].ServerID"), Description: "ServerID is the Server UUID associated with this attachment."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable display name for the volume."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Human-readable description for the volume."},
			{Name: "volume_type", Type: proto.ColumnType_STRING, Description: "The type of volume to create, either SATA or SSD."},
			{Name: "snapshot_id", Type: proto.ColumnType_STRING, Description: "The ID of the snapshot from which the volume was created."},
			{Name: "source_vol_id", Type: proto.ColumnType_STRING, Description: "The ID of another block storage volume from which the current volume was created."},
			{Name: "backup_id", Type: proto.ColumnType_STRING, Description: "The backup ID, from which the volume was restored. This field is supported since 3.47 microversion."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Arbitrary key-value pairs defined by the user."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "UserID is the id of the user who created the volume."},
			{Name: "bootable", Type: proto.ColumnType_BOOL, Description: "Indicates whether this is a bootable volume."},
			{Name: "encrypted", Type: proto.ColumnType_BOOL, Description: "Encrypted denotes if the volume is encrypted."},
			{Name: "replication_status", Type: proto.ColumnType_STRING, Description: "ReplicationStatus is the status of replication."},
			{Name: "consistency_group_id", Type: proto.ColumnType_STRING, Description: "ConsistencyGroupID is the consistency group ID."},
			{Name: "multiattach", Type: proto.ColumnType_BOOL, Description: "Multiattach denotes if the volume is multi-attach capable."},
			{Name: "volume_image_metadata", Type: proto.ColumnType_JSON, Description: "Image metadata entries, only included for volumes that were created from an image, or from a snapshot of a volume originally created from an image."},
		},
	}
}

func listVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_volume.listVolume", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_volume.listVolume", "connection_error", err)
		return nil, err
	}

	// get volumes
	allPages, err := volumes.List(blockStorageClient, volumes.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_volume.listVolume", "query_error", err)
		return nil, err
	}

	allVolumes, err := volumes.ExtractVolumes(allPages)
	for _, volume := range allVolumes {

		// gophercloud requires to extract details for each volume
		volumeDetails, err := volumes.Get(blockStorageClient, volume.ID).Extract()
		if err != nil {
			logger.Error("openstack_volume.listVolume", "query_error", err)
		}
		d.StreamListItem(ctx, volumeDetails)
	}

	return nil, nil
}

func getVolume(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_volume.getVolume", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	// get volume
	volume, err := volumes.Get(blockStorageClient, id).Extract()
	if err != nil {
		logger.Error("openstack_volume.getVolume", "query_error", err)
		return nil, err
	}

	return volume, nil
}
