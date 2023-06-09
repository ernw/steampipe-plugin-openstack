# Table: openstack_fip

A floating IP is an IP address in a public network that can be attached to an instance to make it publicly accessible. Only one floating IP address can be attached to an instance.

## Examples

### Basic floating IP info

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip;
```

### Floating IP by ID

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip
where
  id = 'ccb677bb-d54f-4af9-8127-48ad87c76f74';
```

### Floating IPs created in the last 30 days

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 30;
```

### Floating IPs updated in the last 90 days

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip
where
  date_part('day',current_date::timestamp - updated_at::timestamp) <= 90;
```

### All active floating IPs

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip
where
  status = 'ACTIVE';
```