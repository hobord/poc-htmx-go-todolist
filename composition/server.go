package composition

import (
	"context"
	"log/slog"
	"os"

	"github.com/hobord/poc-htmx-go-todolist/entities"
	"github.com/hobord/poc-htmx-go-todolist/services/logger"
)

type ServerServices struct {
	Log logger.Logger
}

func NewServerServices(ctx context.Context, conf entities.ServerConfig) (*ServerServices, error) {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	return &ServerServices{
		Log: log,
	}, nil
}
