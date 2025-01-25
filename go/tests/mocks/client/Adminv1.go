// Code generated by mockery v2.51.1. DO NOT EDIT.

package client

import (
	adminv1connect "github.com/metal-stack/api/go/admin/v1/adminv1connect"

	mock "github.com/stretchr/testify/mock"
)

// Adminv1 is an autogenerated mock type for the Adminv1 type
type Adminv1 struct {
	mock.Mock
}

// Tenant provides a mock function with no fields
func (_m *Adminv1) Tenant() adminv1connect.TenantServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Tenant")
	}

	var r0 adminv1connect.TenantServiceClient
	if rf, ok := ret.Get(0).(func() adminv1connect.TenantServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(adminv1connect.TenantServiceClient)
		}
	}

	return r0
}

// Token provides a mock function with no fields
func (_m *Adminv1) Token() adminv1connect.TokenServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Token")
	}

	var r0 adminv1connect.TokenServiceClient
	if rf, ok := ret.Get(0).(func() adminv1connect.TokenServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(adminv1connect.TokenServiceClient)
		}
	}

	return r0
}

// NewAdminv1 creates a new instance of Adminv1. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminv1(t interface {
	mock.TestingT
	Cleanup(func())
}) *Adminv1 {
	mock := &Adminv1{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
