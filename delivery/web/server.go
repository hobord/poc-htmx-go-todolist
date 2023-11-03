package web

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type server struct {
	services   *composition.ServerServices
	httpServer http.Server
}

func NewServer(ctx context.Context, conf entities.Config, services *composition.ServerServices) (Server, error) {
	r, err := router.NewRouter(ctx, conf, services)
	if err != nil {
		return nil, err
	}

	return &server{
		services: services,
		httpServer: http.Server{
			Addr:    fmt.Sprintf(":%d", conf.HtttPort),
			Handler: r,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		},
	}, nil
}

func (s *server) Start(ctx context.Context) error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
