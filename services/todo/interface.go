package todo

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Service --structname MockService --output . --outpkg todo --case underscore --filename service_mock.go
type Service interface {
	// GetByID(id string) (*entities.Todo, error)
	GetAll(user string) ([]*entities.Todo, error)
	GetByGroup(user string, group string) ([]*entities.Todo, error)
	GetAllGroup(user string) (map[string][]*entities.Todo, error)
	Create(user, group, title string) (*entities.Todo, error)
	// Update(todo *entities.Todo) error
	// Delete(id string) error
	// SetCompleted(id string, completed bool) error

	// move todo to higher or lower priority and/or move to another group
	// SetPriority(todo, higher, from, to string) error
}
