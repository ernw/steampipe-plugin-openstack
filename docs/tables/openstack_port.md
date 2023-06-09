# Table: openstack_port

A port is a connection point for attaching a single device, such as the NIC of a server, to a network. The port also describes the associated network configuration, such as the MAC and IP addresses to be used on that port.

## Examples

### Basic port info

```sql
select
  name,
  description,
  status,
  id,
  mac_address,
  fixed_ips,
  project_id,
  device_owner,
  security_groups,
  created_at
from
  openstack_port;
```

### Port by ID

```sql
select
  name,
  description,
  status,
  id,
  mac_address,
  fixed_ips,
  project_id,
  device_owner,
  security_groups,
  created_at
from
  openstack_port
where id = '07df44e6-b45b-49e8-b701-303a6f33615d';
```

### Ports created in the last 30 days

```sql
select
  name,
  description,
  id,
  mac_address,
  fixed_ips,
  project_id,
  device_owner,
  security_groups,
  created_at
from
  openstack_port
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 30;
```

### All active ports

```sql
select
  name,
  description,
  id,
  mac_address,
  fixed_ips,
  project_id,
  device_owner,
  security_groups,
  created_at
from
  openstack_port
where
  status = 'ACTIVE';
```