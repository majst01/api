// Code generated by mockery v2.43.2. DO NOT EDIT.

package apiv1connect

import (
	connect "connectrpc.com/connect"
	apiv1 "github.com/metal-stack/api/go/api/v1"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceHandler is an autogenerated mock type for the UserServiceHandler type
type UserServiceHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *UserServiceHandler) Get(_a0 context.Context, _a1 *connect.Request[apiv1.UserServiceGetRequest]) (*connect.Response[apiv1.UserServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv1.UserServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.UserServiceGetRequest]) (*connect.Response[apiv1.UserServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv1.UserServiceGetRequest]) *connect.Response[apiv1.UserServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv1.UserServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv1.UserServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserServiceHandler creates a new instance of UserServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceHandler {
	mock := &UserServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
