# Table: openstack_servergroup

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
  openstack_servergroup;
```