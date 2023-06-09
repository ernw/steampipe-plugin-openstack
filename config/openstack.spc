connection "openstack" {
    plugin    = "ernw/openstack"

    # The HTTP endpoint is REQUIRED to work with the Identity API of the appropriate version.
    # Can also be set with the environment variable "OS_AUTH_URL"
    # identity_endpoint = "http://example.com/identity/v3"

    # Username is REQUIRED if using Identity V2 API.
    # For Identity V3, either UserID or a combination of Username and DomainID or DomainName is REQUIRED.
    # Can also be set with the environment variable "OS_USERNAME" and "OS_USER_ID"
    # username = "admin"
    # user_id = "d8e8fca2dc0f896fd7cb4cb0031ba249"

    # Password is REQUIRED and specifies the password of the user.
    # Can also be set with the environment variable "OS_PASSWORD"
    # password = "changeme"

    # Passcode is OPTIONAL and used in TOTP authentication method.
    # Can also be set with the environment variable "OS_PASSCODE"
    # passcode = "123456"

    # At most one of DomainID and DomainName is REQUIRED if using Username with Identity V3.
    # Otherwise, either are OPTIONAL.
    # Can also be set with the environment variable "OS_DOMAIN_ID" and "OS_DOMAIN_NAME"
    # domain_id = "default"
    # domain_name = "Default"

    # The ProjectId or ProjectName is REQUIRED for Identity V3.
    # Some providers REQUIRE both.
    # Can also be set with the environment variable "OS_PROJECT_ID" and "OS_PROJECT_NAME".
    # project_id = "3e666015f769bf30cda73a1a1e9b794a"
    # project_name = "my_project"

    # AllowReauth should be set to true if you want to cache your credentials 
    # in memory, and to allow attempts to re-authenticate automatically if/when your token
    # expires. This setting is OPTIONAL and defaults to false. 
    # allow_reauth = false

    # Further information: https://docs.openstack.org/python-openstackclient/latest/cli/authentication.html
}
