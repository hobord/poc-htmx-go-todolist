package todo

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Service --structname MockService --inpackage --case underscore --disable-version-string
type Service interface {
	GetTodoGroupsByUserID(userID string) ([]*entities.TodoGroup, error)
	GetTodoGroupByID(todoGroupID string) (*entities.TodoGroup, error)

	CreateTodoGroup(todoGroup *entities.TodoGroup) error
	UpdateTodoGroup(todoGroup *entities.TodoGroup) error
	DeleteTodoGroup(todoGroupID string) error

	AddTodoItem(todoItem *entities.TodoItem) error
	UpdateTodoItem(todoItem *entities.TodoItem) error
	SortTodoItems(ids []string) error
	DeleteTodoItem(todoItemID string) error
	DeleteCompletedTodoItems(groupID string) error
}
