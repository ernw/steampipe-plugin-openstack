---
organization: ERNW
category: ["software development"]
icon_url: "/images/plugins/ernw/openstack.svg"
brand_color: "#ed1944"
display_name: OpenStack
name: openstack
description: Steampipe plugin to query cloud resource information from OpenStack deployments.
og_description: Query OpenStack with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/ernw/openstack-graphic.png"
---

# OpenStack + Steampipe

[OpenStack](https://www.openstack.org/) is the most widely deployed open source cloud software in the world.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

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

## Documentation

- **[Table definitions & examples →](/plugins/ernw/openstack/tables)**

## Get started

### Install

Download and install the latest OpenStack plugin:

```bash
steampipe plugin install ernw/openstack
```

### Credentials

| Item | Description |
| - | - |
| Credentials | [Get credentials from your OpenStack deployment](https://docs.openstack.org/mitaka/cli-reference/common/cli_set_environment_variables_using_openstack_rc.html). |
| Resolution | 1. Credentials explicitly set in a Steampipe config file (`~/.steampipe/config/openstack.spc`).<br /> 2. Credentials specified in environment variables.|

### Configuration

Installing the latest OpenStack plugin will create a config file (`~/.steampipe/config/openstack.spc`) with a single connection named `openstack`.

Configure your account details in `~/.steampipe/config/openstack.spc`:

```hcl
connection "openstack" {
    plugin    = "ernw/openstack"

    # The HTTP endpoint is REQUIRED to work with the Identity API of the appropriate version.
    # Can also be set with the environment variable "OS_AUTH_URL"
    identity_endpoint = "http://example.com/identity/v3"

    # Username is REQUIRED if using Identity V2 API.
    # For Identity V3, either UserID or a combination of Username and DomainID or DomainName is REQUIRED.
    # Can also be set with the environment variable "OS_USERNAME" and "OS_USER_ID"
    username = "admin"
    # user_id = "d8e8fca2dc0f896fd7cb4cb0031ba249"

    # Password is REQUIRED and specifies the password of the user.
    # Can also be set with the environment variable "OS_PASSWORD"
    password = "changeme"

    # Passcode is OPTIONAL and used in TOTP authentication method.
    # Can also be set with the environment variable "OS_PASSCODE"
    # passcode = "123456"

    # At most one of DomainID and DomainName is REQUIRED if using Username with Identity V3.
    # Otherwise, either are OPTIONAL.
    # Can also be set with the environment variable "OS_DOMAIN_ID" and "OS_DOMAIN_NAME"
    domain_id = "default"
    # domain_name = "Default"

    # The ProjectId or ProjectName is REQUIRED for Identity V3.
    # Some providers REQUIRE both.
    # Can also be set with the environment variable "OS_PROJECT_ID" and "OS_PROJECT_NAME".
    project_id = "3e666015f769bf30cda73a1a1e9b794a"
    # project_name = "my_project"

    # AllowReauth should be set to true if you want to cache your credentials 
    # in memory, and to allow attempts to re-authenticate automatically if/when your token
    # expires. This setting is OPTIONAL and defaults to false. 
    # allow_reauth = false

    # Further information: https://docs.openstack.org/python-openstackclient/latest/cli/authentication.html
}
```

Environment variables are also available as an alternate configuration method:

- `OS_AUTH_URL`
- `OS_USERNAME`
- `OS_USER_ID`
- `OS_PASSWORD`
- `OS_PASSCODE`
- `OS_DOMAIN_ID`
- `OS_DOMAIN_NAME`
- `OS_PROJECT_ID`
- `OS_PROJECT_NAME`

## Get involved

* Open source: https://github.com/ernw/steampipe-plugin-openstack
* Community: [Slack Channel](https://steampipe.io/community/join)