// Code generated by mockery v2.51.1. DO NOT EDIT.

package apiv1connect

import (
	connect "connectrpc.com/connect"
	apiv1 "github.com/metal-stack/api/go/api/v1"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PartitionServiceHandler is an autogenerated mock type for the PartitionServiceHandler type
type PartitionServiceHandler struct {
	mock.Mock
}

// Capacity provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceHandler) Capacity(_a0 context.Context, _a1 *connect.Request[apiv1.PartitionServiceCapacityRequest]) (*connect.Response[apiv1.PartitionServiceCapacityResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Capacity")
	}

	var r0 *connect.Response[apiv1.PartitionServiceCapacityResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceCapacityRequest]) (*connect.Response[apiv1.PartitionServiceCapacityResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceCapacityRequest]) *connect.Response[apiv1.PartitionServiceCapacityResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.PartitionServiceCapacityResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.PartitionServiceCapacityRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceHandler) Get(_a0 context.Context, _a1 *connect.Request[apiv1.PartitionServiceGetRequest]) (*connect.Response[apiv1.PartitionServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv1.PartitionServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceGetRequest]) (*connect.Response[apiv1.PartitionServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceGetRequest]) *connect.Response[apiv1.PartitionServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.PartitionServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.PartitionServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *PartitionServiceHandler) List(_a0 context.Context, _a1 *connect.Request[apiv1.PartitionServiceListRequest]) (*connect.Response[apiv1.PartitionServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv1.PartitionServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceListRequest]) (*connect.Response[apiv1.PartitionServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.PartitionServiceListRequest]) *connect.Response[apiv1.PartitionServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.PartitionServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.PartitionServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPartitionServiceHandler creates a new instance of PartitionServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPartitionServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *PartitionServiceHandler {
	mock := &PartitionServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
