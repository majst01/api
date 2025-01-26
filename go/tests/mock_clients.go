// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	"github.com/metal-stack/api/go/admin/v1/adminv1connect"
	"github.com/metal-stack/api/go/api/v1/apiv1connect"
	apiclient "github.com/metal-stack/api/go/client"
	adminv1mocks "github.com/metal-stack/api/go/tests/mocks/admin/v1/adminv1connect"
	apiv1mocks "github.com/metal-stack/api/go/tests/mocks/api/v1/apiv1connect"

	"github.com/stretchr/testify/mock"
)

type (
	client struct {
		adminv1service *adminv1
		apiv1service   *apiv1
	}

	ClientMockFns struct {
		Adminv1Mocks *Adminv1MockFns
		Apiv1Mocks   *Apiv1MockFns
	}

	wrapper struct {
		t *testing.T
	}
	adminv1 struct {
		partitionservice *adminv1mocks.PartitionServiceClient
		tenantservice    *adminv1mocks.TenantServiceClient
		tokenservice     *adminv1mocks.TokenServiceClient
	}

	Adminv1MockFns struct {
		Partition func(m *mock.Mock)
		Tenant    func(m *mock.Mock)
		Token     func(m *mock.Mock)
	}
	apiv1 struct {
		healthservice    *apiv1mocks.HealthServiceClient
		ipservice        *apiv1mocks.IPServiceClient
		methodservice    *apiv1mocks.MethodServiceClient
		partitionservice *apiv1mocks.PartitionServiceClient
		projectservice   *apiv1mocks.ProjectServiceClient
		tenantservice    *apiv1mocks.TenantServiceClient
		tokenservice     *apiv1mocks.TokenServiceClient
		userservice      *apiv1mocks.UserServiceClient
		versionservice   *apiv1mocks.VersionServiceClient
	}

	Apiv1MockFns struct {
		Health    func(m *mock.Mock)
		IP        func(m *mock.Mock)
		Method    func(m *mock.Mock)
		Partition func(m *mock.Mock)
		Project   func(m *mock.Mock)
		Tenant    func(m *mock.Mock)
		Token     func(m *mock.Mock)
		User      func(m *mock.Mock)
		Version   func(m *mock.Mock)
	}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
		adminv1service: w.Adminv1(fns.Adminv1Mocks),
		apiv1service:   w.Apiv1(fns.Apiv1Mocks),
	}
}

func (c *client) Adminv1() apiclient.Adminv1 {
	return c.adminv1service
}
func (c *client) Apiv1() apiclient.Apiv1 {
	return c.apiv1service
}

func (w wrapper) Adminv1(fns *Adminv1MockFns) *adminv1 {
	return newadminv1(w.t, fns)
}

func newadminv1(t *testing.T, fns *Adminv1MockFns) *adminv1 {
	a := &adminv1{
		partitionservice: adminv1mocks.NewPartitionServiceClient(t),
		tenantservice:    adminv1mocks.NewTenantServiceClient(t),
		tokenservice:     adminv1mocks.NewTokenServiceClient(t),
	}

	if fns != nil {
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}

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

func (w wrapper) Apiv1(fns *Apiv1MockFns) *apiv1 {
	return newapiv1(w.t, fns)
}

func newapiv1(t *testing.T, fns *Apiv1MockFns) *apiv1 {
	a := &apiv1{
		healthservice:    apiv1mocks.NewHealthServiceClient(t),
		ipservice:        apiv1mocks.NewIPServiceClient(t),
		methodservice:    apiv1mocks.NewMethodServiceClient(t),
		partitionservice: apiv1mocks.NewPartitionServiceClient(t),
		projectservice:   apiv1mocks.NewProjectServiceClient(t),
		tenantservice:    apiv1mocks.NewTenantServiceClient(t),
		tokenservice:     apiv1mocks.NewTokenServiceClient(t),
		userservice:      apiv1mocks.NewUserServiceClient(t),
		versionservice:   apiv1mocks.NewVersionServiceClient(t),
	}

	if fns != nil {
		if fns.Health != nil {
			fns.Health(&a.healthservice.Mock)
		}
		if fns.IP != nil {
			fns.IP(&a.ipservice.Mock)
		}
		if fns.Method != nil {
			fns.Method(&a.methodservice.Mock)
		}
		if fns.Partition != nil {
			fns.Partition(&a.partitionservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.Token != nil {
			fns.Token(&a.tokenservice.Mock)
		}
		if fns.User != nil {
			fns.User(&a.userservice.Mock)
		}
		if fns.Version != nil {
			fns.Version(&a.versionservice.Mock)
		}

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
