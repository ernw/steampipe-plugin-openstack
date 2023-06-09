# Table: openstack_role_assignment

A 3-tuple combining keystone entities in the form of role-user/group-domain or role-user/group-project. Role assignments are domain specific.

## Examples

### Basic role assignment info

```sql
select
  user_id,
  group_id,
  scope_project_id,
  scope_domain_id,
  scope_role_id
from
  openstack_role_assignment;
```

### All role-group assignemnts

```sql
select
  group_id,
  scope_project_id,
  scope_domain_id,
  scope_role_id
from
  openstack_role_assignment
where
  group_id is not null;
```

### All user-group assignemnts

```sql
select
  user_id,
  scope_project_id,
  scope_domain_id,
  scope_role_id
from
  openstack_role_assignment
where
  user_id is not null;
```

### All user-domain assignments

```sql
select
  user_id,
  scope_project_id,
  scope_domain_id,
  scope_role_id
from
  openstack_role_assignment
where
  scope_domain_id is not null;
```