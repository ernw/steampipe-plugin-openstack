# Table: openstack_aggregate

A host aggregate partitions the hypervisor hosts in an OpenStack cloud. This can be used to partition the hypervisors, e.g., depending on their hardware capabilities.

## Examples

### Basic aggregate info

```sql
select
  availability_zone,
  hosts,
  id,
  name
from
  openstack_aggregate;
```

### Aggregate by ID

```sql
select
  availability_zone,
  hosts,
  id,
  name
from
  openstack_aggregate
where
  id = 1;
```

### All deleted aggregates

```sql
select
  availability_zone,
  hosts,
  hosts,
  metadata,
  name,
  created_at,
  updated_at,
  deleted_at
from
  openstack_aggregate
where
  deleted = true;
```

### Aggregates created in the last 90 days

```sql
select
  availability_zone,
  hosts,
  id,
  name
from
  openstack_aggregate
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 90;
```