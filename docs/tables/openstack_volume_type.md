# Table: openstack_volume_type

Volume types allow specifying attributes for volumes.

## Examples

### Basic volume type info

```sql
select
  id,
  name,
  description,
  extra_specs,
  is_public,
  public_access
from
  openstack_volume_type;
```

### Volume type by ID

```sql
select
  id,
  name,
  description,
  extra_specs,
  is_public,
  public_access
from
  openstack_volume_type
where
  id = '78d3a6e6-56cc-4835-adf6-e5f676e1a362';
```


### All public volume types

```sql
select
  id,
  name,
  description,
  extra_specs,
  is_public,
  public_access
from
  openstack_volume_type
where
  is_public = true;
```