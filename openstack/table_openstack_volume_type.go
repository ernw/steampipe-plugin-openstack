package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackVolumeType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_volume_type",
		Description: "Table of all volume types.",
		List: &plugin.ListConfig{
			Hydrate: listVolumeType,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolumeType,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the volume type."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable display name for the volume type."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Human-readable description for the volume type."},
			{Name: "extra_specs", Type: proto.ColumnType_JSON, Description: "Arbitrary key-value pairs defined by the user."},
			{Name: "is_public", Type: proto.ColumnType_BOOL, Description: "Whether the volume type is publicly visible."},
			{Name: "qos_specs_id", Type: proto.ColumnType_STRING, Description: "Qos Spec ID."},
			{Name: "public_access", Type: proto.ColumnType_BOOL, Description: "Volume Type access public attribute."},
		},
	}
}

func listVolumeType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_volume_type.listVolumeType", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_volume_type.listVolumeType", "connection_error", err)
		return nil, err
	}

	// get volumetypes
	allPages, err := volumetypes.List(blockStorageClient, volumetypes.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_volume_type.listVolumeType", "query_error", err)
		return nil, err
	}

	allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	for _, volumeType := range allVolumeTypes {
		d.StreamListItem(ctx, volumeType)
	}

	return nil, nil
}

func getVolumeType(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_volume_type.getVolumeType", "connection_error", err)
		return nil, err
	}

	// get block storage client from provider
	blockStorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})

	// get volume type
	volumeType, err := volumetypes.Get(blockStorageClient, id).Extract()
	if err != nil {
		logger.Error("openstack_volume_type.getVolumeType", "query_error", err)
		return nil, err
	}

	return volumeType, nil
}
