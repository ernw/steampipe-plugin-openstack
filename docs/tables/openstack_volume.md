# Table: openstack_volume

A volume is a virtual disk that holds data. Volumes can be attached to instances.

## Examples

### Basic volume info

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume;
```

### Volume by ID

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  id = 'f52d0d12-99b9-4c91-8f21-04b2a5b12a1e';
```

### All bootable volumes bigger than 10 GB

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  size > 10;
```

### All bootable volumes

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  bootable = true;
```

### All encrypted volumes

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  encrypted = true;
```

### Volumes created in the last 30 days

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  date_part('day',current_date::timestamp - created_at::timestamp) <= 30;
```

### Volumes not updated in the last 90 days

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  date_part('day',current_date::timestamp - updated_at::timestamp) >= 90;
```