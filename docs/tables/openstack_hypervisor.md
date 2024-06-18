# Table: openstack_server

A hypervisor isolates the hypervisor operating system from the virtual machines (servers).

## Examples

### Basic hypervisor info

```sql
select
  id,
  hypervisor_hostname,
  host_ip,
  state,
  hypervisor_type
from
  openstack_hypervisor;
```

## All running hypervisors

```sql
select
  id,
  hypervisor_hostname,
  host_ip,
  state,
  hypervisor_type
from
  openstack_hypervisor
where
  state = 'up';
```

## System resource usage of running hypervisors

```sql
select
  id,
  hypervisor_hostname,
  host_ip,
  vcpus,
  vcpus_used,
  disk_available_least,
  free_disk_gb,
  free_ram_mb,
  local_gb,
  local_gb_used,
  running_vms,
  memory_mb,
  memory_mb_used
from
  openstack_hypervisor
where
  state = 'up';
```

## All disabled hypervisors

```sql
select
  id,
  hypervisor_hostname,
  host_ip,
  state,
  hypervisor_type
from
  openstack_hypervisor
where
  status = 'disabled';
```