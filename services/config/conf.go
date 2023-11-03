package config

import "github.com/hobord/poc-htmx-go-todolist/entities"

func NewConfig() (*entities.Config, error) {
	return &entities.Config{
		HtttPort: 8080,
	}, nil
}
