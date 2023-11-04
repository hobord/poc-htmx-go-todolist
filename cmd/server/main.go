package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/dal/config/viper"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web"
	"github.com/hobord/poc-htmx-go-todolist/services/config"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	v := viper.NewReader()
	config := config.NewService(v)

	conf, err := config.GetServerConfig()
	if err != nil {
		return err
	}

	services, err := composition.NewServerServices(ctx, conf)
	if err != nil {
		return err
	}

	log := services.Log

	webServer, err := web.NewServer(ctx, conf, services)
	if err != nil {
		log.Error("Could not create web server", "error", err)
		return err
	}

	err = webServer.Start(ctx)
	if err != nil {
		log.Error("Could not start web server", "error", err)
		return err
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	err = webServer.Stop(ctx)
	if err != nil {
		log.Error("Could not stop web server", "error", err)
		return err
	}

	log.Info("Server stopped")

	return nil
}
