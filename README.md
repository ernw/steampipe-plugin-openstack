# OpenStack Plugin for Steampipe

Use SQL to query cloud resources and their configuration from [OpenStack](https://www.openstack.org/).

* **[Get started →](https://hub.steampipe.io/plugins/ernw/openstack)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/ernw/openstack/tables)
* Community: [Slack Channel](https://steampipe.io/community/join)
* Get involved: [Issues](https://github.com/ernw/steampipe-plugin-openstack/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ernw/openstack
```

Configure your [credentials](https://hub.steampipe.io/plugins/ernw/openstack#credentials) and [config file](https://hub.steampipe.io/plugins/ernw/openstack#configuration).

Configure your account details in `~/.steampipe/config/openstack.spc`:

```hcl
connection "openstack" {
    plugin    = "ernw/openstack"

    # Authentication information
    identity_endpoint = "http://example.com/identity/v3"
    username = "admin"
    password = "changeme"
    domain_id = "default"
    project_id = "3e666015f769bf30cda73a1a1e9b794a"
}
```

Environment variables are also available as an alternate configuration method:

```bash
export OS_AUTH_URL="http://example.com/identity/v3"
export OS_USERNAME="admin"
export OS_PASSWORD="changeme"
export OS_DOMAIN_ID="default"
export OS_PROJECT_ID="3e666015f769bf30cda73a1a1e9b794a"
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  name,
  description,
  email,
  enabled
from
  openstack_user;
```

```
+-------------------+-----------------------------+---------+
| name              | email                       | enabled |
+-------------------+-----------------------------+---------+
| demo              | demo@example.com            | true    |
| admin             | admin@testproject.com       | true    |
| reader            | reader@testproject.com      | true    |
+-------------------+-----------------------------+---------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/ernw/steampipe-plugin-openstack.git
cd steampipe-plugin-openstack
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/openstack.spc
```

Try it!

```
steampipe query
> .inspect openstack
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/ernw/steampipe-plugin-openstack/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [OpenStack Plugin](https://github.com/ernw/steampipe-plugin-openstack/labels/help%20wanted)