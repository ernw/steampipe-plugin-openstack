# Table: openstack_secgroup

A security group is a container of security group rules. A default security group is assigned to newly created instances.

## Examples

### Basic security group info

```sql
select
  name,
  description,
  rules,
  id,
  project_id,
  updated_at,
  created_at
from
  openstack_secgroup;
```

### All security groups that have been updated

```sql
select
  name,
  description,
  rules,
  id,
  project_id,
  updated_at,
  created_at
from
  openstack_secgroup
where
  updated_at != created_at;
```