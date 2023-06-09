package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableOpenstackSecurityGroupRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_security_group_rule",
		Description: "Table of all security group rules.",
		List: &plugin.ListConfig{
			Hydrate: listSecurityGroupRule,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSecurityGroupRule,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The UUID for the security group."},
			{Name: "direction", Type: proto.ColumnType_STRING, Description: "The direction in which the security group rule is applied. The only values allowed are 'ingress' or 'egress'."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the rule."},
			{Name: "ether_type", Type: proto.ColumnType_STRING, Description: "Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress or egress rules."},
			{Name: "sec_group_id", Type: proto.ColumnType_STRING, Description: "The security group ID to associate with this security group rule."},
			{Name: "port_range_min", Type: proto.ColumnType_INT, Description: "The minimum port number in the range that is matched by the security group rule."},
			{Name: "port_range_max", Type: proto.ColumnType_INT, Description: "The maximum port number in the range that is matched by the security group rule."},
			{Name: "protocol", Type: proto.ColumnType_STRING, Description: "The protocol that is matched by the security group rule. Valid values are 'tcp', 'udp', 'icmp' or an empty string."},
			{Name: "remote_group_id", Type: proto.ColumnType_STRING, Description: "The remote group ID to be associated with this security group rule. You can specify either RemoteGroupID or RemoteIPPrefix."},
			{Name: "remote_ip_prefix", Type: proto.ColumnType_STRING, Description: "The remote IP prefix to be associated with this security group rule. You can specify either RemoteGroupID or RemoteIPPrefix. This attribute matches the specified IP prefix as the source IP address of the IP packet."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "TenantID is the project owner of this security group rule."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "ProjectID is the project owner of this security group rule."},
		},
	}
}

func listSecurityGroupRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_security_group_rule.listSecurityGroupRule", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_security_group_rule.listSecurityGroupRule", "connection_error", err)
		return nil, err
	}

	// get security group rules
	allPages, err := rules.List(networkClient, rules.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_security_group_rule.listSecurityGroupRule", "query_error", err)
		return nil, err
	}

	allRules, err := rules.ExtractRules(allPages)
	for _, rule := range allRules {
		d.StreamListItem(ctx, rule)
	}

	return nil, nil
}

func getSecurityGroupRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_security_group_rule.getSecurityGroupRule", "connection_error", err)
		return nil, err
	}

	// get network client from provider
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	// get security group rule
	rule, err := rules.Get(networkClient, id).Extract()
	if err != nil {
		logger.Error("openstack_security_group_rule.getSecurityGroupRule", "query_error", err)
		return nil, err
	}

	return rule, nil
}
