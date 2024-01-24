package api

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/justinas/alice"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/health"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/index"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/todo"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/middleware"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func CreateHandler(
	ctx context.Context,
	conf entities.ServerConfig,
	services *composition.ServerServices,
	assets fs.FS,
) (http.Handler, error) {
	log := services.Log

	log.Debug("Create http routing handlers")

	api := router.NewRouter()

	// static assets
	log.Debug("Register handler", "method", http.MethodGet, "path", "/assets/*path")
	api.Handle("/assets/*", http.FileServer(http.FS(assets)))

	// health check
	log.Debug("Register handler", "method", http.MethodGet, "path", "/health")

	if services.HealthService != nil {
		healthCheck, err := health.NewCheck(services.HealthService)
		if err != nil {
			return nil, err
		}

		api.MethodFunc(http.MethodGet, "/health", healthCheck.Health)
	}

	// root
	{
		// empty prefix for root level
		root := router.NewGroup(api, "").
			WithMiddlewares(
				middleware.WithLogger(services.Log),
			)

		// index
		indexHandler, err := index.NewHandler(services.TodoService)
		if err != nil {
			return nil, err
		}

		root.GET("/", indexHandler.IndexPage)

		// api.Handler(http.MethodGet, "/",
		// 	alice.New(middleware.Logger).
		// 		Then(router.HandlerFunc(indexHandler.IndexPage)),
		// )

		// todos
		{
			todoHandler, err := todo.NewHandler(services.TodoService)
			if err != nil {
				return nil, err
			}

			todoGroup := root.Group("/group")
			todoGroup.GET("/", indexHandler.IndexPage)

			// create todo group
			todoGroup.POST("/", todoHandler.CreateTodoGroup)

			// delete todo group
			todoGroup.DELETE("/{groupID}", todoHandler.DeleteTodoGroup)

			// sort group items
			todoGroup.POST("/{groupID}/sort", todoHandler.SortItems)

			todoItem := root.Group("/todo")

			// add todo item
			todoItem.POST("/", todoHandler.CreateItem)

			// delete todo item
			todoItem.DELETE("/{itemID}", todoHandler.DeleteItem)
		}

		// test
		{
			testGroup := root.Group("/test")
			testGroup.GET("/page", indexHandler.IndexPage)

			{
				sub := testGroup.Group("/sub") // .WithMiddlewares(middleware.Logger)
				sub.GET("/page", indexHandler.IndexPage)
			}
		}

	}

	handler := alice.New(middleware.PanicRecovery).Then(api)

	return handler, nil
}
