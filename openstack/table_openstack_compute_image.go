package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackComputeImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_compute_image",
		Description: "Table of all images.",
		List: &plugin.ListConfig{
			Hydrate: listComputeImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getComputeImage,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of an image."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Created contain ISO-8601 timestamp of when the image was created."},
			{Name: "min_disk", Type: proto.ColumnType_INT, Description: "MinDisk is the minimum amount of disk a flavor must have to be able to create a server based on the image, measured in GB."},
			{Name: "min_ram", Type: proto.ColumnType_INT, Transform: transform.FromField("MinRAM"), Description: "MinRAM is the minimum amount of RAM a flavor must have to be able to create a server based on the image, measured in MB."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name provides a human-readable moniker for the OS image."},
			{Name: "progress", Type: proto.ColumnType_INT, Description: "The Progress and Status fields indicate image-creation status."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status is the current status of the image."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "Updated contain ISO-8601 timestamp of when the state of the image last changed."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata provides free-form key/value pairs that further describe the image."},
		},
	}
}

func listComputeImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_compute_image.listComputeImage", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_compute_image.listComputeImage", "connection_error", err)
		return nil, err
	}

	// get images
	allPages, err := images.ListDetail(computeClient, images.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_compute_image.listComputeImage", "query_error", err)
		return nil, err
	}

	allImages, err := images.ExtractImages(allPages)
	for _, image := range allImages {
		d.StreamListItem(ctx, image)
	}

	return nil, nil
}

func getComputeImage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_compute_image.getComputeImage", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	// get image
	image, err := images.Get(computeClient, id).Extract()
	if err != nil {
		logger.Error("openstack_compute_image.getComputeImage", "query_error", err)
		return nil, err
	}

	return image, nil
}
