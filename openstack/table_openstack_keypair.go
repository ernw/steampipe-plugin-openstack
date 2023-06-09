package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackKeypair(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_keypair",
		Description: "Table of all key pairs.",
		List: &plugin.ListConfig{
			Hydrate: listKeypair,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getKeypair,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is used to refer to this keypair from other services within this region."},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "Fingerprint is a short sequence of bytes that can be used to authenticate or validate a longer public key."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Description: "PublicKey is the public key from this pair, in OpenSSH format."},
			{Name: "private_key", Type: proto.ColumnType_STRING, Description: "PrivateKey is the private key from this pair, in PEM format."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "UserID is the user who owns this KeyPair."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the keypair."},
		},
	}
}

func listKeypair(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_keypair.listKeypair", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_keypair.listKeypair", "connection_error", err)
		return nil, err
	}

	// get keypairs
	allPages, err := keypairs.List(computeClient, keypairs.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_keypair.listKeypair", "query_error", err)
		return nil, err
	}

	allKeypairs, err := keypairs.ExtractKeyPairs(allPages)
	for _, keypair := range allKeypairs {
		d.StreamListItem(ctx, keypair)
	}

	return nil, nil
}

func getKeypair(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	name := d.EqualsQuals["name"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_keypair.getKeypair", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	// get keypair
	keypair, err := keypairs.Get(computeClient, name, keypairs.GetOpts{}).Extract()
	if err != nil {
		logger.Error("openstack_keypair.getKeypair", "query_error", err)
		return nil, err
	}

	return keypair, nil
}
