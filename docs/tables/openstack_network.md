# Table: openstack_network

A network contains one or more subnets.

## Examples

### Basic network info

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network;
```

### Network by ID 

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  id = '4d33bcfe-215d-44e9-8986-93033a20789f';
```

### All active networks

```sql
select
  name,
  description,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  status = 'ACTIVE';
```

### All shared networks

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  shared = true;
```

### All networks created in the last 30 days

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 30;
```

### All networks updated in the last 90 days

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  date_part('day',current_date::timestamp - updated_at::timestamp) <= 90;
```