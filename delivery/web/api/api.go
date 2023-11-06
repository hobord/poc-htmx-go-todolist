package api

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/health"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/index"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/middleware"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func CreateHandler(ctx context.Context, conf entities.ServerConfig, services *composition.ServerServices, assets fs.FS) (http.Handler, error) {
	api := httprouter.New()

	// static assets
	api.Handler(http.MethodGet, "/assets/*path", http.FileServer(http.FS(assets)))

	// health check
	healthCheck := health.NewCheck(services.HealthService)
	api.HandlerFunc(http.MethodGet, "/health", healthCheck.Health)

	// empty prefix for root level
	root := router.NewGroup(api, "").
		WithMiddlewares(middleware.WithLogger(services.Log))

	// index
	indexHandler := index.NewHandler(services.TodoService)
	root.GET("/", indexHandler.IndexPage)

	// api.Handler(http.MethodGet, "/",
	// 	alice.New(middleware.Logger).
	// 		Then(router.HandlerFunc(indexHandler.IndexPage)),
	// )

	testGroup := root.Group("/test")

	testGroup.GET("/page", indexHandler.IndexPage)

	sub := testGroup.Group("/sub") // .WithMiddlewares(middleware.Logger)
	sub.GET("/page", indexHandler.IndexPage)

	r := alice.New(middleware.PanicRecovery).Then(api)

	return r, nil
}
