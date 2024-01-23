// Code generated by mockery v2.40.1. DO NOT EDIT.

package config

import (
	entities "github.com/hobord/poc-htmx-go-todolist/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

// ReadServerConfig provides a mock function with given fields:
func (_m *MockReader) ReadServerConfig() (entities.ServerConfig, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadServerConfig")
	}

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

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
