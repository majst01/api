// Code generated by mockery v2.51.1. DO NOT EDIT.

package apiv1connect

import (
	connect "connectrpc.com/connect"
	apiv1 "github.com/metal-stack/api/go/api/v1"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TokenServiceClient is an autogenerated mock type for the TokenServiceClient type
type TokenServiceClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceClient) Create(_a0 context.Context, _a1 *connect.Request[apiv1.TokenServiceCreateRequest]) (*connect.Response[apiv1.TokenServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[apiv1.TokenServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceCreateRequest]) (*connect.Response[apiv1.TokenServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceCreateRequest]) *connect.Response[apiv1.TokenServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.TokenServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.TokenServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceClient) List(_a0 context.Context, _a1 *connect.Request[apiv1.TokenServiceListRequest]) (*connect.Response[apiv1.TokenServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv1.TokenServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceListRequest]) (*connect.Response[apiv1.TokenServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceListRequest]) *connect.Response[apiv1.TokenServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.TokenServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.TokenServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: _a0, _a1
func (_m *TokenServiceClient) Revoke(_a0 context.Context, _a1 *connect.Request[apiv1.TokenServiceRevokeRequest]) (*connect.Response[apiv1.TokenServiceRevokeResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Revoke")
	}

	var r0 *connect.Response[apiv1.TokenServiceRevokeResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceRevokeRequest]) (*connect.Response[apiv1.TokenServiceRevokeResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.TokenServiceRevokeRequest]) *connect.Response[apiv1.TokenServiceRevokeResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.TokenServiceRevokeResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.TokenServiceRevokeRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenServiceClient creates a new instance of TokenServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenServiceClient {
	mock := &TokenServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
