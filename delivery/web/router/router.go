package router

import (
	"context"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/health"
	"github.com/hobord/poc-htmx-go-todolist/entities"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctx context.Context, conf entities.Config, services *composition.ServerServices) (http.Handler, error) {
	router := httprouter.New()

	// health check
	healthCheck := health.NewCheck()
	router.HandlerFunc(http.MethodGet, "/health", healthCheck.Health)

	return router, nil
}
