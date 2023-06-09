# Table: openstack_compute_image

An image contains a virtual disk that holds a bootable operating system.

## Examples

### Basic image info

```sql
select
  name,
  status,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image;
```

### Image by ID

```sql
select
  name,
  status,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  id = '63a81077-d72f-4429-95b0-897f975af068';
```

### All active images

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  status = 'ACTIVE';
```

### All images created in the last 30 days

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  date_part('day',current_date::timestamp - created::timestamp) <= 30;
```

### All images not updated in the last 90 days

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  date_part('day',current_date::timestamp - updated::timestamp) >= 90;
```

### All images not yet built

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  progress < 100;
```