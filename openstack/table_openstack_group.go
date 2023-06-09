package openstack

import (
	"context"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenstackGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_group",
		Description: "Table of all groups.",
		List: &plugin.ListConfig{
			Hydrate: listGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGroup,
		},
		Columns: []*plugin.Column{
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.Description"), Description: "Description describes the group purpose."},
			{Name: "domain_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.DomainID"), Description: "DomainID is the domain ID the group belongs to."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.ID"), Description: "ID is the unique ID of the group."},
			{Name: "group_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.Name"), Description: "Name is the name of the group."},
			{Name: "extra", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group.Extra"), Description: "Extra is a collection of miscellaneous key/values."},
			{Name: "member_list", Type: proto.ColumnType_JSON, Description: "UserIDs of all members."},
		},
	}
}

type GroupEntry struct {
	Group      groups.Group
	MemberList string
}

func listGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_group.listGroup", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logger.Error("openstack_group.listGroup", "connection_error", err)
		return nil, err
	}

	// get groups
	allPages, err := groups.List(identityClient, groups.ListOpts{}).AllPages()
	allGroups, err := groups.ExtractGroups(allPages)
	if err != nil {
		logger.Error("openstack_group.listGroup", "query_error", err)
		return nil, err
	}

	for _, group := range allGroups {

		// get users of groups
		var memberList []string
		allPages, err = users.ListInGroup(identityClient, group.ID, users.ListOpts{}).AllPages()
		allMembers, err := users.ExtractUsers(allPages)
		if err != nil {
			logger.Error("openstack_group.listGroup", "query_error", err)
			return nil, err
		}

		for _, member := range allMembers {
			memberList = append(memberList, member.ID)
		}

		membersJSON, err := json.Marshal(memberList)
		d.StreamListItem(ctx, GroupEntry{group, string(membersJSON)})
	}

	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_group.getGroup", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logger.Error("openstack_group.getGroup", "connection_error", err)
		return nil, err
	}

	// get group
	group, err := groups.Get(identityClient, id).Extract()
	if err != nil {
		logger.Error("openstack_group.getGroup", "query_error", err)
		return nil, err
	}

	// get users of group
	var memberList []string
	allPages, err := users.ListInGroup(identityClient, group.ID, users.ListOpts{}).AllPages()
	allMembers, err := users.ExtractUsers(allPages)
	if err != nil {
		logger.Error("openstack_group.getGroup", "query_error", err)
		return nil, err
	}

	for _, member := range allMembers {
		memberList = append(memberList, member.ID)
	}

	membersJSON, err := json.Marshal(memberList)
	return GroupEntry{*group, string(membersJSON)}, nil
}
