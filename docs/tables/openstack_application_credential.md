# Table: openstack_application_credential

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
  openstack_application_credential;
```

### Application credential by ID and UserID

```sql
select
  name,
  description,
  unrestricted,
  project_id,
  roles,
  expires_at
from
  openstack_application_credential
where
  id = 'e9b1ae13d4254d0c9fb9387f64e4a953' and user_id='6c85c7ff83f24ac0a64f56db18782ecb';
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
  openstack_application_credential
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
  openstack_application_credential
where
  expires_at < cast(now() as date);
```