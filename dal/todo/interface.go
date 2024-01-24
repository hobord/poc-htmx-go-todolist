package todo

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Reader --structname MockReader  --inpackage --case underscore --disable-version-string
type Reader interface {
	GetTodoGroupsByUserID(userID string) ([]*entities.TodoGroup, error)
	GetTodoGroupByID(todoGroupID string) (*entities.TodoGroup, error)
	GetTodoItemByID(id string) (*entities.TodoItem, error)
}

//go:generate mockery --name Writer --structname MockWriter  --inpackage --case underscore --disable-version-string
type Writer interface {
	WriteTodoGroup(todoGroup *entities.TodoGroup) error
	DeleteTodoGroup(todoGroupID string) error
	WriteTodoItem(todoItem *entities.TodoItem) error
	DeleteTodoItem(todoItemID string) error
}

//go:generate mockery --name ReaderWriter --structname MockReaderWriter  --inpackage --case underscore --disable-version-string
type ReaderWriter interface {
	Reader
	Writer
}
