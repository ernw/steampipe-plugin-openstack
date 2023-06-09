# Table: openstack_security_group_rule

A security group rule allows ingress or egress traffic to and from certain instances to a specified IP address or IP range.

## Examples

### Basic security group rule info

```sql
select
  id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_security_group_rule;
```

### Security group rule by ID

```sql
select
  id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_security_group_rule
where
  id = 'ef7884e6-9ee7-4547-932f-a65b4d1d5e8e';
```

### All security group rules with TCP protocol

```sql
select
  id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_security_group_rule
where
  protocol = 'tcp';
```

### All IPv4 security group rules

```sql
select
  id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_security_group_rule
where
  ether_type = 'IPv4';
```

### All security group rules with specified port range

```sql
select
  id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_security_group_rule
where
  port_range_min is not null;
```