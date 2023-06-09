# Table: openstack_server

A server is a VM that runs on a hypervisor. A server is launched from an image.

## Examples

### Basic server info

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server;
```

### Server by ID

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  id = '544c2a99-fb22-4810-b08b-f9bf2eb6f991';
```

### All servers not created in the last 30 days

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  date_part('day',current_date::timestamp - created::timestamp) <= 30;
```

### All servers not updated in the last 90 days

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  date_part('day',current_date::timestamp - updated::timestamp) >= 90;
```

### All shutoff servers

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  status = 'SHUTOFF';
```

### All shelved servers

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  status = 'SHELVED_OFFLOADED';
```