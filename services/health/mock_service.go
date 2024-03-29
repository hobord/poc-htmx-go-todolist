// Code generated by mockery. DO NOT EDIT.

package health

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddChecker provides a mock function with given fields: _a0
func (_m *MockService) AddChecker(_a0 func() error) {
	_m.Called(_a0)
}

// Health provides a mock function with given fields:
func (_m *MockService) Health() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Health")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
