package config

import "github.com/hobord/poc-htmx-go-todolist/entities"

//go:generate mockery --name Service --structname MockService --inpackage --case underscore --disable-version-string
type Service interface {
	GetServerConfig() (entities.ServerConfig, error)
}
