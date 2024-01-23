// Code generated by mockery. DO NOT EDIT.

package todo

import (
	entities "github.com/hobord/poc-htmx-go-todolist/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockWriter is an autogenerated mock type for the Writer type
type MockWriter struct {
	mock.Mock
}

// DeleteTodoGroup provides a mock function with given fields: todoGroupID
func (_m *MockWriter) DeleteTodoGroup(todoGroupID string) error {
	ret := _m.Called(todoGroupID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodoGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(todoGroupID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodoItem provides a mock function with given fields: todoItemID
func (_m *MockWriter) DeleteTodoItem(todoItemID string) error {
	ret := _m.Called(todoItemID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodoItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(todoItemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteTodoGroup provides a mock function with given fields: todoGroup
func (_m *MockWriter) WriteTodoGroup(todoGroup *entities.TodoGroup) error {
	ret := _m.Called(todoGroup)

	if len(ret) == 0 {
		panic("no return value specified for WriteTodoGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.TodoGroup) error); ok {
		r0 = rf(todoGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteTodoItem provides a mock function with given fields: todoItem
func (_m *MockWriter) WriteTodoItem(todoItem *entities.TodoItem) error {
	ret := _m.Called(todoItem)

	if len(ret) == 0 {
		panic("no return value specified for WriteTodoItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.TodoItem) error); ok {
		r0 = rf(todoItem)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockWriter creates a new instance of MockWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriter {
	mock := &MockWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
