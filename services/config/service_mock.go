// Code generated by mockery v2.36.0. DO NOT EDIT.

package config

import (
	entities "github.com/hobord/poc-htmx-go-todolist/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// GetServerConfig provides a mock function with given fields:
func (_m *MockService) GetServerConfig() (entities.ServerConfig, error) {
	ret := _m.Called()

	var r0 entities.ServerConfig
	var r1 error
	if rf, ok := ret.Get(0).(func() (entities.ServerConfig, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() entities.ServerConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entities.ServerConfig)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
