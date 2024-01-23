package user

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Reader --structname MockReader  --inpackage --case underscore --disable-version-string
type Reader interface {
	GetUserByID(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}

//go:generate mockery --name Writer --structname MockWriter  --inpackage --case underscore --disable-version-string
type Writer interface {
	AddUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
}

//go:generate mockery --name ReaderWriter --structname MockReaderWriter  --inpackage --case underscore --disable-version-string
type ReaderWriter interface {
	Reader
	Writer
}
