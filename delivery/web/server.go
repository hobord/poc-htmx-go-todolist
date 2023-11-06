package web

import (
	"context"
	"embed"
	"fmt"
	"net"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/api"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

var (
	//go:embed all:assets/*
	assetsFS embed.FS
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type server struct {
	services   *composition.ServerServices
	httpServer http.Server
}

func NewServer(ctx context.Context, conf entities.ServerConfig, services *composition.ServerServices) (Server, error) {
	apiHandler, err := api.CreateHandler(ctx, conf, services, assetsFS)
	if err != nil {
		return nil, err
	}

	return &server{
		services: services,
		httpServer: http.Server{
			Addr:    fmt.Sprintf(":%d", conf.HttpPort),
			Handler: apiHandler,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		},
	}, nil
}

func (s *server) Start(ctx context.Context) error {
	s.services.Log.Info("Starting server", "httpPort", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	s.services.Log.Info("Stopping server")

	return s.httpServer.Shutdown(ctx)
}
