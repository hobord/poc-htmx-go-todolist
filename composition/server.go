package composition

import (
	"context"
	"log/slog"
	"os"

	"github.com/hobord/poc-htmx-go-todolist/entities"
	"github.com/hobord/poc-htmx-go-todolist/services/health"
	"github.com/hobord/poc-htmx-go-todolist/services/logger"
	"github.com/hobord/poc-htmx-go-todolist/services/todo"
)

type ServerServices struct {
	Log           logger.Logger
	HealthService health.Service
	TodoService   todo.Service
}

func NewServerServices(ctx context.Context, conf entities.ServerConfig) (*ServerServices, error) {
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return &ServerServices{
		Log:           log,
		HealthService: createMockHealtService(),
		TodoService:   createMockTodoService(),
	}, nil
}
