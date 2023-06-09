# Table: openstack_secgrouprule

A security group rule allows ingress or egress traffic to and from certain instances to a specified IP address or IP range.

## Examples

### Basic security group rule info

```sql
select
  rule_id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_secgrouprule;
```

### All security group rules with TCP protocol

```sql
select
  rule_id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_secgrouprule
where
  protocol = 'tcp';
```

### All IPv4 security group rules

```sql
select
  rule_id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_secgrouprule
where
  ether_type = 'IPv4';
```

### All security group rules with specified port range

```sql
select
  rule_id,
  description,
  ether_type,
  port_range_min,
  port_range_max,
  protocol,
  remote_ip_prefix
from
  openstack_secgrouprule
where
  port_range_min is not null;
```