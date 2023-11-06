package todo

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Reader --structname MockReader --output . --outpkg todo --case underscore --filename reader_mock.go
type Reader interface {
	GetByID(id string) (*entities.Todo, error)
	GetAll(user string) ([]*entities.Todo, error)
	GetByGroup(user string, group string) ([]*entities.Todo, error)
}

//go:generate mockery --name Writer --structname MockWriter --output . --outpkg todo --case underscore --filename writer_mock.go
type Writer interface {
	Create(todo *entities.Todo) error
	Update(todo *entities.Todo) error
	Delete(id string) error
	SetCompleted(id string, completed bool) error
	SetPriority(todos []*entities.Todo) error
}

//go:generate mockery --name ReaderWriter --structname MockReaderWriter --output . --outpkg todo --case underscore --filename reader_writer_mock.go
type ReaderWriter interface {
	Reader
	Writer
}
