package config

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Reader --structname MockReader --output . --outpkg config --case underscore --filename reader_mock.go
type Reader interface {
	ReadServerConfig() (entities.ServerConfig, error)
}
