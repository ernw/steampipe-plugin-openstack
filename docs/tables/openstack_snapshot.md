# Table: openstack_snapshot

A snapshot is the backup of a volume.

## Examples

### Basic snapshot info

```sql
select
  id,
  name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot;
```

### Snapshot by ID

```sql
select
  id,
  name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  id = '91f4432d-5da5-475d-84b9-c68789dddb70';
```

### All snapshots bigger than 1 GB

```sql
select
  id,
  name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  size > 1;
```

### All available snapshots

```sql
select
  id,
  name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  status = 'available';
```

### All snapshots not updated in the last 14 days

```sql
select
  id,
  name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  date_part('day',current_date::timestamp - updated_at::timestamp) >= 14;
```
