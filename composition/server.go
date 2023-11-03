package composition

import (
	"context"

	"github.com/hobord/poc-htmx-go-todolist/entities"
)

type ServerServices struct {
}

func NewServerServices(ctx context.Context, conf entities.Config) (*ServerServices, error) {
	return &ServerServices{}, nil
}
