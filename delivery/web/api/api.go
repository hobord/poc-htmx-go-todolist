package api

import (
	"context"
	"io/fs"
	"net/http"

	"github.com/hobord/routegroup"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/health"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/index"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/handler/todo"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/middleware"
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

	mux := http.NewServeMux()

	root := routegroup.NewGroup(
		routegroup.WithMux(mux),
		routegroup.WithRegisterCallback(routegroup.RegisterRouteCallback),
		routegroup.WithRegisterPanicCallback(routegroup.RegisterPanicHandler),
		routegroup.WithMiddlewares(routegroup.Recover),
	)

	// static assets
	log.Debug("Register handler", "method", http.MethodGet, "path", "/assets/*path")
	root.Handle("GET /assets/*", http.FileServer(http.FS(assets)))

	// health check
	log.Debug("Register handler", "method", http.MethodGet, "path", "/health")

	if services.HealthService != nil {
		healthCheck, err := health.NewCheck(services.HealthService)
		if err != nil {
			return nil, err
		}

		root.HandleFunc("GET /health/{$}", healthCheck.Health)
	}

	root.Use(middleware.Logger)

	// index
	indexHandler, err := index.NewHandler(services.TodoService)
	if err != nil {
		return nil, err
	}

	root.HandleFunc("GET /{$}", indexHandler.IndexPage)

	// todos
	{
		todoHandler, err := todo.NewHandler(services.TodoService)
		if err != nil {
			return nil, err
		}

		// todo group
		{
			todoGroup := root.SubGroup("/group")
			todoGroup.HandleFunc("GET /{$}", indexHandler.IndexPage)

			// create todo group
			todoGroup.HandleFunc("POST /{$}", todoHandler.CreateTodoGroup)

			// delete todo group
			todoGroup.HandleFunc("DELETE /{groupID}", todoHandler.DeleteTodoGroup)

			// sort group items
			todoGroup.HandleFunc("POST /{groupID}/sort", todoHandler.SortItems)
		}

		// todo item
		{
			todoItem := root.SubGroup("/todo")

			// add todo item
			todoItem.HandleFunc("POST /{$}", todoHandler.CreateItem)

			// delete todo item
			todoItem.HandleFunc("DELETE /{itemID}", todoHandler.DeleteItem)
		}
	}

	// // test
	// {
	// 	testGroup := root.SubGroup("/test")
	// 	testGroup.HandleFunc("GET /page", indexHandler.IndexPage)

	// 	{
	// 		sub := testGroup.SubGroup("/sub") // .WithMiddlewares(middleware.Logger)
	// 		sub.HandleFunc("GET /page", indexHandler.IndexPage)
	// 	}
	// }

	return mux, nil
}
