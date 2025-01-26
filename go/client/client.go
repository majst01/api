// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	compress "github.com/klauspost/connect-compress/v2"

	"github.com/metal-stack/api/go/admin/v1/adminv1connect"
	"github.com/metal-stack/api/go/api/v1/apiv1connect"
)

type (
	Client interface {
		Adminv1() Adminv1
		Apiv1() Apiv1
	}
	client struct {
		config DialConfig
	}
	Adminv1 interface {
		Partition() adminv1connect.PartitionServiceClient
		Tenant() adminv1connect.TenantServiceClient
		Token() adminv1connect.TokenServiceClient
	}

	adminv1 struct {
		partitionservice adminv1connect.PartitionServiceClient
		tenantservice    adminv1connect.TenantServiceClient
		tokenservice     adminv1connect.TokenServiceClient
	}

	Apiv1 interface {
		Health() apiv1connect.HealthServiceClient
		IP() apiv1connect.IPServiceClient
		Method() apiv1connect.MethodServiceClient
		Partition() apiv1connect.PartitionServiceClient
		Project() apiv1connect.ProjectServiceClient
		Tenant() apiv1connect.TenantServiceClient
		Token() apiv1connect.TokenServiceClient
		User() apiv1connect.UserServiceClient
		Version() apiv1connect.VersionServiceClient
	}

	apiv1 struct {
		healthservice    apiv1connect.HealthServiceClient
		ipservice        apiv1connect.IPServiceClient
		methodservice    apiv1connect.MethodServiceClient
		partitionservice apiv1connect.PartitionServiceClient
		projectservice   apiv1connect.ProjectServiceClient
		tenantservice    apiv1connect.TenantServiceClient
		tokenservice     apiv1connect.TokenServiceClient
		userservice      apiv1connect.UserServiceClient
		versionservice   apiv1connect.VersionServiceClient
	}
)

func New(config DialConfig) Client {
	return &client{
		config: config,
	}
}

func (c client) Adminv1() Adminv1 {
	a := &adminv1{
		partitionservice: adminv1connect.NewPartitionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: adminv1connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tokenservice: adminv1connect.NewTokenServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *adminv1) Partition() adminv1connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *adminv1) Tenant() adminv1connect.TenantServiceClient {
	return c.tenantservice
}
func (c *adminv1) Token() adminv1connect.TokenServiceClient {
	return c.tokenservice
}

func (c client) Apiv1() Apiv1 {
	a := &apiv1{
		healthservice: apiv1connect.NewHealthServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		ipservice: apiv1connect.NewIPServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		methodservice: apiv1connect.NewMethodServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		partitionservice: apiv1connect.NewPartitionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		projectservice: apiv1connect.NewProjectServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: apiv1connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		tokenservice: apiv1connect.NewTokenServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		userservice: apiv1connect.NewUserServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
		versionservice: apiv1connect.NewVersionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			compress.WithAll(compress.LevelBalanced),
		),
	}
	return a
}

func (c *apiv1) Health() apiv1connect.HealthServiceClient {
	return c.healthservice
}
func (c *apiv1) IP() apiv1connect.IPServiceClient {
	return c.ipservice
}
func (c *apiv1) Method() apiv1connect.MethodServiceClient {
	return c.methodservice
}
func (c *apiv1) Partition() apiv1connect.PartitionServiceClient {
	return c.partitionservice
}
func (c *apiv1) Project() apiv1connect.ProjectServiceClient {
	return c.projectservice
}
func (c *apiv1) Tenant() apiv1connect.TenantServiceClient {
	return c.tenantservice
}
func (c *apiv1) Token() apiv1connect.TokenServiceClient {
	return c.tokenservice
}
func (c *apiv1) User() apiv1connect.UserServiceClient {
	return c.userservice
}
func (c *apiv1) Version() apiv1connect.VersionServiceClient {
	return c.versionservice
}
