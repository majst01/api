// Code generated by mockery v2.52.1. DO NOT EDIT.

package client

import (
	adminv2connect "github.com/metal-stack/api/go/metalstack/admin/v2/adminv2connect"

	mock "github.com/stretchr/testify/mock"
)

// Adminv2 is an autogenerated mock type for the Adminv2 type
type Adminv2 struct {
	mock.Mock
}

// Partition provides a mock function with no fields
func (_m *Adminv2) Partition() adminv2connect.PartitionServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Partition")
	}

	var r0 adminv2connect.PartitionServiceClient
	if rf, ok := ret.Get(0).(func() adminv2connect.PartitionServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(adminv2connect.PartitionServiceClient)
		}
	}

	return r0
}

// Tenant provides a mock function with no fields
func (_m *Adminv2) Tenant() adminv2connect.TenantServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Tenant")
	}

	var r0 adminv2connect.TenantServiceClient
	if rf, ok := ret.Get(0).(func() adminv2connect.TenantServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(adminv2connect.TenantServiceClient)
		}
	}

	return r0
}

// Token provides a mock function with no fields
func (_m *Adminv2) Token() adminv2connect.TokenServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Token")
	}

	var r0 adminv2connect.TokenServiceClient
	if rf, ok := ret.Get(0).(func() adminv2connect.TokenServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(adminv2connect.TokenServiceClient)
		}
	}

	return r0
}

// NewAdminv2 creates a new instance of Adminv2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminv2(t interface {
	mock.TestingT
	Cleanup(func())
}) *Adminv2 {
	mock := &Adminv2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
