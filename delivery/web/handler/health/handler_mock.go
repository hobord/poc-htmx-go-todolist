// Code generated by mockery v2.40.1. DO NOT EDIT.

package health

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

// Health provides a mock function with given fields: w, r
func (_m *MockHandler) Health(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
