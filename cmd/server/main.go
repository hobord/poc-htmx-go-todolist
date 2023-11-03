package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web"
	"github.com/hobord/poc-htmx-go-todolist/services/config"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	services, err := composition.NewServerServices(ctx, *conf)
	if err != nil {
		panic(err)
	}

	webServer, err := web.NewServer(ctx, *conf, services)
	if err != nil {
		panic(err)
	}

	err = webServer.Start(ctx)
	if err != nil {
		panic(err)
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	err = webServer.Stop(ctx)
	if err != nil {
		panic(err)
	}
}
