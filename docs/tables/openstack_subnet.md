# Table: openstack_subnet

A subnet has a block of IP addresses and allows network communication between instances. Instances can be assigned to subnets.

## Examples

### Basic subnet info

```sql
select
  name,
  description,
  id,
  ip_version,
  cidr,
  gateway_ip,
  allocation_pools,
  enable_dhcp,
  project_id
from
  openstack_subnet;
```

### Subnet by ID

```sql
select
  name,
  description,
  id,
  ip_version,
  cidr,
  gateway_ip,
  allocation_pools,
  enable_dhcp,
  project_id
from
  openstack_subnet
where
  id = '2e4a84c9-5ee2-42d2-9608-4194b0d5e865';
```

### All IPv4 subnets

```sql
select
  name,
  description,
  id,
  ip_version,
  cidr,
  gateway_ip,
  allocation_pools,
  enable_dhcp,
  project_id
from
  openstack_subnet
where
  ip_version = '4';
```

### All subnets with DHCP

```sql
select
  name,
  description,
  id,
  ip_version,
  cidr,
  gateway_ip,
  allocation_pools,
  enable_dhcp,
  project_id
from
  openstack_subnet
where
  enable_dhcp = true;
```