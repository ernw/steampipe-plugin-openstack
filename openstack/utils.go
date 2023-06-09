package openstack

import (
	"context"
	"os"
	"reflect"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*gophercloud.ProviderClient, error) {

	logger := plugin.Logger(ctx)

	// check cache for connection
	cacheKey := "openstack"
	if cacheData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cacheData.(*gophercloud.ProviderClient), nil
	}

	// get connection properties from env var or from config file if env var not present
	openstackConfig := GetConfig(d.Connection)

	var envIsPresent bool
	opts := gophercloud.AuthOptions{}

	opts.IdentityEndpoint, envIsPresent = os.LookupEnv("OS_AUTH_URL")
	if !envIsPresent && openstackConfig.IdentityEndpoint != nil {
		opts.IdentityEndpoint = *openstackConfig.IdentityEndpoint
	}

	opts.Username, envIsPresent = os.LookupEnv("OS_USERNAME")
	if !envIsPresent && openstackConfig.Username != nil {
		opts.Username = *openstackConfig.Username
	}

	opts.UserID, envIsPresent = os.LookupEnv("OS_USER_ID")
	if !envIsPresent && openstackConfig.UserID != nil {
		opts.UserID = *openstackConfig.UserID
	}

	opts.Password, envIsPresent = os.LookupEnv("OS_PASSWORD")
	if !envIsPresent && openstackConfig.Password != nil {
		opts.Password = *openstackConfig.Password
	}

	opts.Passcode, envIsPresent = os.LookupEnv("OS_PASSCODE")
	if !envIsPresent && openstackConfig.Passcode != nil {
		opts.Passcode = *openstackConfig.Passcode
	}

	opts.DomainID, envIsPresent = os.LookupEnv("OS_DOMAIN_ID")
	if !envIsPresent && openstackConfig.DomainID != nil {
		opts.DomainID = *openstackConfig.DomainID
	}

	opts.DomainName, envIsPresent = os.LookupEnv("OS_DOMAIN_NAME")
	if !envIsPresent && openstackConfig.DomainName != nil {
		opts.DomainName = *openstackConfig.DomainName
	}

	opts.TenantID, envIsPresent = os.LookupEnv("OS_PROJECT_ID")
	if !envIsPresent && openstackConfig.ProjectID != nil {
		opts.TenantID = *openstackConfig.ProjectID
	}

	opts.TenantName, envIsPresent = os.LookupEnv("OS_PROJECT_NAME")
	if !envIsPresent && openstackConfig.ProjectName != nil {
		opts.TenantName = *openstackConfig.ProjectName
	}

	// default to false
	if openstackConfig.AllowReauth != nil {
		opts.AllowReauth = *openstackConfig.AllowReauth
	}

	// get ProviderClient struct
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logger.Error("Connection properties were invalid or uncomplete", err)
		return nil, err
	}

	// save to cache
	d.ConnectionManager.Cache.Set(cacheKey, provider)

	return provider, nil
}

func getMapItemByKey(ctx context.Context, d *transform.TransformData) (interface{}, error) {

	v := reflect.ValueOf(d.Value)
	p := reflect.ValueOf(d.Param)

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			if key.Interface() == p.Interface() {
				return v.MapIndex(key), nil
			}
		}
	}

	return nil, nil
}
