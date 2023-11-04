package config

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Service --structname MockService --output . --outpkg config --case underscore --filename service_mock.go
type Service interface {
	GetServerConfig() (entities.ServerConfig, error)
}
