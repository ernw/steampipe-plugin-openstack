# Table: openstack_security_group

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
  openstack_security_group;
```

### Security group by ID

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
  openstack_security_group
where
  id = 'b29d3c6b-1d21-4094-8d75-1fc0dac7f1af';
```

### All security groups created in the last 30 days

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
  openstack_security_group
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 30;
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
  openstack_security_group
where
  updated_at != created_at;
```