# Table: openstack_group

A group represents a container for one or multiple users. A group can only be part of one domain.

## Examples

### Basic group info

```sql
select
  group_name,
  description,
  domain_id,
  member_list,
  id
from
  openstack_group;
```

### Group by ID

```sql
select
  group_name,
  description,
  domain_id,
  member_list,
  id
from
  openstack_group
where
  id = '47e4a9c0f1574ed99a3d5e9a3d91dd80';
```

### All groups in default domain

```sql
select
  group_name,
  description,
  domain_id,
  member_list,
  id
from
  openstack_group
where
  domain_id = 'default';
```