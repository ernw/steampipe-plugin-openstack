package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackSubnet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_subnet",
		Description: "Table of all subnets.",
		List: &plugin.ListConfig{
			Hydrate: listSubnet,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSubnet,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "UUID representing the subnet."},
			{Name: "network_id", Type: proto.ColumnType_STRING, Description: "UUID of the parent network."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable name for the subnet. Might not be unique."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the subnet."},
			{Name: "ip_version", Type: proto.ColumnType_INT, Description: "IP version, either `4' or `6'."},
			{Name: "cidr", Type: proto.ColumnType_STRING, Transform: transform.FromField("CIDR"), Description: "CIDR representing IP range for this subnet, based on IP version."},
			{Name: "gateway_ip", Type: proto.ColumnType_STRING, Description: "Default gateway used by devices in this subnet."},
			{Name: "dns_nameservers", Type: proto.ColumnType_JSON, Transform: transform.FromField("DNSNameservers"), Description: "DNS name servers used by hosts in this subnet."},
			{Name: "allocation_pools", Type: proto.ColumnType_JSON, Description: "Sub-ranges of CIDR available for dynamic allocation to ports."},
			{Name: "host_routes", Type: proto.ColumnType_JSON, Description: "Routes that should be used by devices with IPs from this subnet (not including local subnet route)."},
			{Name: "enable_dhcp", Type: proto.ColumnType_BOOL, Transform: transform.FromField("EnableDHCP"), Description: "Specifies whether DHCP is enabled for this subnet or not."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the subnet."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the subnet."},
			{Name: "ipv6_address_mode", Type: proto.ColumnType_STRING, Transform: transform.FromField("IPv6AddressMode"), Description: "The IPv6 address modes specifies mechanisms for assigning IPv6 IP addresses."},
			{Name: "ipv6_ra_mode", Type: proto.ColumnType_STRING, Transform: transform.FromField("IPv6RAMode"), Description: "The IPv6 router advertisement specifies whether the networking service should transmit ICMPv6 packets."},
			{Name: "subnet_pool_id", Type: proto.ColumnType_STRING, Description: "SubnetPoolID is the id of the subnet pool associated with the subnet."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags optionally set via extensions/attributestags."},
			{Name: "revision_number", Type: proto.ColumnType_INT, Description: "RevisionNumber optionally set via extensions/standard-attr-revisions."},
		},
	}
}

func listSubnet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_subnet.listSubnet", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_subnet.listSubnet", "connection_error", err)
		return nil, err
	}

	// get subnets
	allPages, err := subnets.List(networkClient, subnets.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_subnet.listSubnet", "query_error", err)
		return nil, err
	}

	allSubnets, err := subnets.ExtractSubnets(allPages)
	for _, subnet := range allSubnets {
		d.StreamListItem(ctx, subnet)
	}

	return nil, nil
}

func getSubnet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_subnet.getSubnet", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get subnet
	subnet, err := subnets.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_subnet.getSubnet", "query_error", err)
		return nil, err
	}

	return subnet, nil
}
