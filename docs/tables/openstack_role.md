# Table: openstack_role

A role grants authorization to an end-user and dictates the authorization level assigned to a user.

## Examples

### Basic role info

```sql
select
  id,
  name,
  extra
from
  openstack_role;
```

### Role by ID

```sql
select
  id,
  name,
  extra
from
  openstack_role
where
  id = 'd104344e6db0467689a3e721edc9dc2b';
```

### Roles with a description

```sql
select
  id,
  name,
  extra
from
  openstack_role
where
  extra -> 'description' != 'null';
```