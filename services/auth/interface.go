package auth

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Service --structname MockService --inpackage --case underscore --disable-version-string
type Service interface {
	RegisterUser(*entities.User) error
	LoginUser(email, password string) (*entities.User, error)
}
