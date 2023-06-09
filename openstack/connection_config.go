package openstack

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type openstackConfig struct {
	IdentityEndpoint *string `cty:"identity_endpoint"`
	Username         *string `cty:"username"`
	UserID           *string `cty:"user_id"`
	Password         *string `cty:"password"`
	Passcode         *string `cty:"passcode"`
	DomainID         *string `cty:"domain_id"`
	DomainName       *string `cty:"domain_name"`
	ProjectID        *string `cty:"project_id"`
	ProjectName      *string `cty:"project_name"`
	AllowReauth      *bool   `cty:"allow_reauth"`
	// TODO: add authentication with application credentials
}

var ConfigSchema = map[string]*schema.Attribute{
	"identity_endpoint": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"user_id": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"passcode": {
		Type: schema.TypeString,
	},
	"domain_id": {
		Type: schema.TypeString,
	},
	"domain_name": {
		Type: schema.TypeString,
	},
	"project_id": {
		Type: schema.TypeString,
	},
	"project_name": {
		Type: schema.TypeString,
	},
	"allow_reauth": {
		Type: schema.TypeBool,
	},
}

func ConfigInstance() interface{} {
	return &openstackConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) openstackConfig {
	if connection == nil || connection.Config == nil {
		return openstackConfig{}
	}
	config, _ := connection.Config.(openstackConfig)

	return config
}
