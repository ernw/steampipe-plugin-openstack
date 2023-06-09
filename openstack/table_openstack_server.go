package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_server",
		Description: "Table of all server instances.",
		List: &plugin.ListConfig{
			Hydrate: listServer,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getServer,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID uniquely identifies this server."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID identifies the tenant owning this server resource."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "UserID uniquely identifies the user account owning the tenant."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name contains the human-readable name for the server."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "Updated contain ISO-8601 timestamp of when the state of the server last changed."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Created contain ISO-8601 timestamp of when the server was created."},
			{Name: "host_id", Type: proto.ColumnType_STRING, Description: "HostID is the host where the server is located in the cloud."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status contains the current operational status of the server, such as IN_PROGRESS or ACTIVE."},
			{Name: "progress", Type: proto.ColumnType_INT, Description: "Progress ranges from 0..100. A request made against the server completes only once Progress reaches 100."},
			{Name: "access_ipv4", Type: proto.ColumnType_STRING, Transform: transform.FromField("AccessIPv4"), Description: "AccessIPv4 contains the IP address of the server."},
			{Name: "access_ipv6", Type: proto.ColumnType_STRING, Transform: transform.FromField("AccessIPv6"), Description: "AccessIPv6 contains the IP address of the server."},
			{Name: "image", Type: proto.ColumnType_JSON, Description: "Image refers to a JSON object, which itself indicates the OS image used to deploy the server."},
			{Name: "flavor", Type: proto.ColumnType_JSON, Description: "Flavor refers to a JSON object, which itself indicates the hardware configuration of the deployed server."},
			{Name: "addresses", Type: proto.ColumnType_JSON, Description: "Addresses includes a list of all IP addresses assigned to the server, keyed by pool."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata includes a list of all user-specified key-value pairs attached to the server."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links includes HTTP references to the itself, useful for passing along to other APIs that might want a server reference."},
			{Name: "key_name", Type: proto.ColumnType_STRING, Description: "KeyName indicates which public key was injected into the server on launch."},
			{Name: "admin_pass", Type: proto.ColumnType_STRING, Description: "AdminPass will generally be empty (\"\").  However, it will contain the administrative password chosen when provisioning a new server without a set AdminPass setting in the first place. Note that this is the ONLY time this field will be valid."},
			{Name: "security_groups", Type: proto.ColumnType_JSON, Description: "SecurityGroups includes the security groups that this instance has applied to it."},
			{Name: "attached_volumes", Type: proto.ColumnType_STRING, Description: "AttachedVolumes includes the volume attachments of this instance."},
			{Name: "fault", Type: proto.ColumnType_STRING, Description: "Fault contains failure information about a server."},
			{Name: "tags", Type: proto.ColumnType_STRING, Description: "Tags is a slice/list of string tags in a server. The requires microversion 2.26 or later."},
			{Name: "server_groups", Type: proto.ColumnType_STRING, Description: "ServerGroups is a slice of strings containing the UUIDs of the server groups to which the server belongs. Currently this can contain at most one entry. New in microversion 2.71."},
		},
	}
}

func listServer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_server.listServer", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_server.listServer", "connection_error", err)
		return nil, err
	}

	// get servers
	allPages, err := servers.List(computeClient, servers.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_server.listServer", "query_error", err)
		return nil, err
	}

	allServers, err := servers.ExtractServers(allPages)
	for _, server := range allServers {
		d.StreamListItem(ctx, server)
	}

	return nil, nil
}

func getServer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_server.getServer", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	// get server
	server, err := servers.Get(computeClient, id).Extract()
	if err != nil {
		logger.Error("openstack_server.getServer", "query_error", err)
		return nil, err
	}

	return server, nil
}
