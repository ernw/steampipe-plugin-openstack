# Table: openstack_availability_zone

Availability Zones are an end-user visible logical abstraction for partitioning a cloud without knowing the physical infrastructure. They can be used to partition a cloud on arbitrary factors, such as location (country, datacenter, rack), network layout and/or power source.

## Examples

### Basic availability zone info

```sql
select
  hosts,
  zone_name,
  zone_state
from
  openstack_availability_zone;
```

### All available availability zones

```sql
select
  hosts,
  zone_name,
  zone_state
from
  openstack_availability_zone
where
  zone_state ->> 'available' = 'true';
```