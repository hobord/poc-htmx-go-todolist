// Code generated by mockery. DO NOT EDIT.

package todo

import (
	entities "github.com/hobord/poc-htmx-go-todolist/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

// GetTodoGroupByID provides a mock function with given fields: todoGroupID
func (_m *MockReader) GetTodoGroupByID(todoGroupID string) (*entities.TodoGroup, error) {
	ret := _m.Called(todoGroupID)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoGroupByID")
	}

	var r0 *entities.TodoGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.TodoGroup, error)); ok {
		return rf(todoGroupID)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.TodoGroup); ok {
		r0 = rf(todoGroupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TodoGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(todoGroupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoGroupsByUserID provides a mock function with given fields: userID
func (_m *MockReader) GetTodoGroupsByUserID(userID string) ([]*entities.TodoGroup, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoGroupsByUserID")
	}

	var r0 []*entities.TodoGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*entities.TodoGroup, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []*entities.TodoGroup); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.TodoGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoItemByID provides a mock function with given fields: id
func (_m *MockReader) GetTodoItemByID(id string) (*entities.TodoItem, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoItemByID")
	}

	var r0 *entities.TodoItem
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.TodoItem, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.TodoItem); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TodoItem)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
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
