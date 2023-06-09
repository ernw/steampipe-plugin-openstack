# Table: openstack_applicationcredential

An application credential consists of a credential ID and a secret string. It is used by applications to interact with OpenStack and deployed resources.

## Examples

### Basic application credential info

```sql
select
  name,
  description,
  unrestricted,
  project_id,
  roles,
  expires_at
from
  openstack_applicationcredential;
```

### All application credentials allowing unrestricted access

```sql
select
  id,
  name,
  description,
  project_id,
  roles,
  expires_at
from
  openstack_applicationcredential
where
  unrestricted = true;
```

### All expired application credentials

```sql
select
  id,
  name,
  description,
  project_id,
  roles
from
  openstack_applicationcredential
where
  expires_at < cast(now() as date);
```