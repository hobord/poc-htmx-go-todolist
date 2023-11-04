package router

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/health"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/index"
	"github.com/hobord/poc-htmx-go-todolist/entities"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctx context.Context, conf entities.ServerConfig, services *composition.ServerServices, assets fs.FS) (http.Handler, error) {
	router := httprouter.New()

	// static assets
	router.Handler(http.MethodGet, "/assets/*path", http.FileServer(http.FS(assets)))

	// health check
	healthCheck := health.NewCheck()
	router.HandlerFunc(http.MethodGet, "/health", healthCheck.Health)

	// index
	indexHandler := index.NewHandler()
	router.HandlerFunc(http.MethodGet, "/", indexHandler.IndexPage)

	return router, nil
}
