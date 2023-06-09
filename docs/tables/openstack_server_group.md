# Table: openstack_server_group

A server group constitutes a container holding one or more instances. Server groups can be used to organize instances depending on their task or properties.

## Examples

### Basic server group info

```sql
select
  name,
  id,
  policies,
  members
from
  openstack_server_group;
```

### Server group by ID

```sql
select
  name,
  id,
  policies,
  members
from
  openstack_server_group;
```