package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud/openstack"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenstackHypervisor(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_hypervisor",
		Description: "Table of all hypervisors.",
		List: &plugin.ListConfig{
			Hydrate: listHypervisor,
			KeyColumns: plugin.KeyColumnSlice{
				&plugin.KeyColumn{
					Name:    "id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getHypervisor,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the hypervisor."},
			{Name: "hypervisor_hostname", Type: proto.ColumnType_STRING, Description: "The hostname of the hypervisor."},
			{Name: "host_ip", Type: proto.ColumnType_STRING, Description: "The hypervisor's IP address.", Transform: transform.FromField("HostIP")},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the hypervisor, either 'up'or 'down'."},
			{Name: "hypervisor_version", Type: proto.ColumnType_STRING, Description: "The version of the hypervisor."},
			{Name: "vcpus", Type: proto.ColumnType_INT, Description: "The total number of vcpus on the hypervisor.", Transform: transform.FromField("VCPUs")},
			{Name: "cpu_vendor", Type: proto.ColumnType_STRING, Description: "The vendor of the CPU.", Transform: transform.FromField("CPUInfo.Vendor")},
			{Name: "cpu_arch", Type: proto.ColumnType_STRING, Description: "The arch of the CPU.", Transform: transform.FromField("CPUInfo.Arch")},
			{Name: "cpu_model", Type: proto.ColumnType_STRING, Description: "The model of the CPU.", Transform: transform.FromField("CPUInfo.Model")},
			{Name: "vcpus_used", Type: proto.ColumnType_INT, Description: "The number of used vcpus on the hypervisor.", Transform: transform.FromField("VCPUsUsed")},
			{Name: "current_workload", Type: proto.ColumnType_INT, Description: "The number of tasks the hypervisor is responsible for."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the hypervisor, either 'enabled' or 'disabled'."},
			{Name: "disk_available_least", Type: proto.ColumnType_INT, Description: "The actual free disk on the hypervisor, measured in GB."},
			{Name: "free_disk_gb", Type: proto.ColumnType_INT, Description: "The free disk remaining on the hypervisor, measured in GB.", Transform: transform.FromField("FreeDiskGB")},
			{Name: "free_ram_mb", Type: proto.ColumnType_INT, Description: "The free RAM in the hypervisor, measured in MB.", Transform: transform.FromField("FreeRamMB")},
			{Name: "local_gb", Type: proto.ColumnType_INT, Description: "The disk space in the hypervisor, measured in GB.", Transform: transform.FromField("LocalGB")},
			{Name: "local_gb_used", Type: proto.ColumnType_INT, Description: "The used disk space of the hypervisor, measured in GB.", Transform: transform.FromField("LocalGBUsed")},
			{Name: "running_vms", Type: proto.ColumnType_INT, Description: "The number of running vms on the hypervisor.", Transform: transform.FromField("RunningVMs")},
			{Name: "memory_mb", Type: proto.ColumnType_INT, Description: "The total memory of the hypervisor, measured in MB.", Transform: transform.FromField("MemoryMB")},
			{Name: "memory_mb_used", Type: proto.ColumnType_INT, Description: "The used memory of the hypervisor, measured in MB.", Transform: transform.FromField("MemoryMBUsed")},
			{Name: "service", Type: proto.ColumnType_STRING, Description: "The service this hypervisor represents."},
			{Name: "hypervisor_type", Type: proto.ColumnType_STRING, Description: "The type of hypervisor."},
		},
	}
}

func listHypervisor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_hypervisor.listHypervisor", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	endpointOpts := getEndpointOpts(d)
	computeClient, err := openstack.NewComputeV2(provider, endpointOpts)
	if err != nil {
		logger.Error("openstack_hypervisor.listHypervisor", "connection_error", err)
		return nil, err
	}

	// get hypervisors
	allPages, err := hypervisors.List(computeClient, hypervisors.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_hypervisor.listHypervisor", "query_error", err)
		return nil, err
	}

	allHypervisors, err := hypervisors.ExtractHypervisors(allPages)
	for _, hypervisor := range allHypervisors {
		d.StreamListItem(ctx, hypervisor)
	}

	return nil, nil
}

func getHypervisor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	id := d.EqualsQuals["id"].GetStringValue()

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_hypervisor.getHypervisor", "connection_error", err)
		return nil, err
	}

	// get compute client from provider
	endpointOpts := getEndpointOpts(d)
	computeClient, err := openstack.NewComputeV2(provider, endpointOpts)

	// get hypervisor
	hypervisor, err := hypervisors.Get(computeClient, id).Extract()
	if err != nil {
		logger.Error("openstack_hypervisor.getHypervisor", "query_error", err)
		return nil, err
	}

	return hypervisor, nil
}
