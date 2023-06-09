package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackPort(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_port",
		Description: "Table of all ports.",
		List: &plugin.ListConfig{
			Hydrate: listPort,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getPort,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "UUID for the port."},
			{Name: "network_id", Type: proto.ColumnType_STRING, Description: "Network that this port is associated with."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Human-readable name for the port. Might not be unique."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Describes the port."},
			{Name: "admin_state_up", Type: proto.ColumnType_BOOL, Description: "Administrative state of port. If false (down), port does not forward packets."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Indicates whether network is currently operational. Possible values include `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional values."},
			{Name: "mac_address", Type: proto.ColumnType_STRING, Transform: transform.FromField("MACAddress"), Description: "Mac address to use on this port."},
			{Name: "fixed_ips", Type: proto.ColumnType_JSON, Transform: transform.FromField("FixedIPs"), Description: "Specifies IP addresses for the port thus associating the port itself with the subnets where the IP addresses are picked from."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of the port."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of the port."},
			{Name: "device_owner", Type: proto.ColumnType_STRING, Description: "Identifies the entity (e.g.: dhcp agent) using this port."},
			{Name: "security_groups", Type: proto.ColumnType_JSON, Description: "Specifies the IDs of any security groups associated with a port."},
			{Name: "device_id", Type: proto.ColumnType_STRING, Description: "Identifies the device (e.g., virtual server) using this port."},
			{Name: "allowed_address_pairs", Type: proto.ColumnType_JSON, Description: "Identifies the list of IP addresses the port will recognize/accept."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags optionally set via extensions/attributestags."},
			{Name: "revision_number", Type: proto.ColumnType_INT, Description: "RevisionNumber optionally set via extensions/standard-attr-revisions."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "UpdatedAt contains an ISO-8601 timestamp of when the state of the port last changed."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "CreatedAt contains an ISO-8601 timestamp of when the port was created."},
		},
	}
}

func listPort(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_port.listPort", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_port.listPort", "connection_error", err)
		return nil, err
	}

	// get ports
	allPages, err := ports.List(networkClient, ports.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_port.listPort", "query_error", err)
		return nil, err
	}

	allPorts, err := ports.ExtractPorts(allPages)
	for _, port := range allPorts {
		d.StreamListItem(ctx, port)
	}

	return nil, nil
}

func getPort(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_port.getPort", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get port
	port, err := ports.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_port.getPort", "query_error", err)
		return nil, err
	}

	return port, nil
}
