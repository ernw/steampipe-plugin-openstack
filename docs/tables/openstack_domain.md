# Table: openstack_domain

A domain is a container for projects, users, and groups. 

## Examples

### Basic domain info

```sql
select
  name,
  description,
  enabled,
  id
from
  openstack_domain;
```

### Domain by ID

```sql
select
  name,
  description,
  id
from
  openstack_domain
where
  id = 'default';
```

### All active domains

```sql
select
  name,
  description,
  id
from
  openstack_domain
where
  enabled = true;
```