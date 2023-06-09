# Table: openstack_user

A user is an individual API consumer who can authenticate and access cloud resources. A user is part of only one domain and the username must be unique to the domain.

## Examples

### Basic user info

```sql
select
  name,
  description,
  email,
  enabled,
  lock_password,
  domain_id,
  password_expires_at,
  default_project_id
from
  openstack_user;
```

### User by ID

```sql
select
  name,
  description,
  email,
  enabled,
  lock_password,
  domain_id,
  password_expires_at,
  default_project_id
from
  openstack_user
where
  id = 'e021d695cc604acdb8866686b51f6321';
```

### All disabled users

```sql
select
  name,
  description,
  email,
  enabled,
  lock_password,
  domain_id,
  password_expires_at,
  default_project_id
from
  openstack_user
where
  enabled = false;
```

### All users with no password expiry date

```sql
select
  name,
  description,
  email,
  enabled,
  lock_password,
  domain_id,
  password_expires_at,
  default_project_id
from
  openstack_user
where
  password_expires_at is null;
```