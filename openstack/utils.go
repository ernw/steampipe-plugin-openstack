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
	cacheKey := "connection"
	if cacheData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cacheData.(*gophercloud.ProviderClient), nil
	}

	// get connection properties from env var or from config file if env var not present
	openstackConfig := GetConfig(d.Connection)

	opts := gophercloud.AuthOptions{}

	opts.IdentityEndpoint = os.Getenv("OS_AUTH_URL")
	if openstackConfig.IdentityEndpoint != nil {
		opts.IdentityEndpoint = *openstackConfig.IdentityEndpoint
	}

	opts.Username = os.Getenv("OS_USERNAME")
	if openstackConfig.Username != nil {
		opts.Username = *openstackConfig.Username
	}

	opts.UserID = os.Getenv("OS_USER_ID")
	if openstackConfig.UserID != nil {
		opts.UserID = *openstackConfig.UserID
	}

	opts.Password = os.Getenv("OS_PASSWORD")
	if openstackConfig.Password != nil {
		opts.Password = *openstackConfig.Password
	}

	opts.Passcode = os.Getenv("OS_PASSCODE")
	if openstackConfig.Passcode != nil {
		opts.Passcode = *openstackConfig.Passcode
	}

	opts.DomainID = os.Getenv("OS_DOMAIN_ID")
	if openstackConfig.DomainID != nil {
		opts.DomainID = *openstackConfig.DomainID
	}

	opts.DomainName = os.Getenv("OS_DOMAIN_NAME")
	if openstackConfig.DomainName != nil {
		opts.DomainName = *openstackConfig.DomainName
	}

	opts.TenantID = os.Getenv("OS_PROJECT_ID")
	if openstackConfig.ProjectID != nil {
		opts.TenantID = *openstackConfig.ProjectID
	}

	opts.TenantName = os.Getenv("OS_PROJECT_NAME")
	if openstackConfig.ProjectName != nil {
		opts.TenantName = *openstackConfig.ProjectName
	}

	// default to false
	if openstackConfig.AllowReauth != nil {
		opts.AllowReauth = *openstackConfig.AllowReauth
	}

	// get ProviderClient struct
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logger.Error("Connection properties are invalid or incomplete", err)
		return nil, err
	}

	// save to cache
	d.ConnectionManager.Cache.Set(cacheKey, provider)

	return provider, nil
}

func getEndpointOpts(d *plugin.QueryData) gophercloud.EndpointOpts {

	// check cache for endpointOpts
	cacheKey := "endpointOpts"
	if cacheData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cacheData.(gophercloud.EndpointOpts)
	}

	endpointOpts := gophercloud.EndpointOpts{}

	// get region from env var or config file
	endpointOpts.Region = os.Getenv("OS_REGION")

	openstackConfig := GetConfig(d.Connection)
	if openstackConfig.Region != nil {
		endpointOpts.Region = *openstackConfig.Region
	}

	// save to cache
	d.ConnectionManager.Cache.Set(cacheKey, endpointOpts)

	return endpointOpts
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
