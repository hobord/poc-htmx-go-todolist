package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/hobord/poc-htmx-go-todolist/composition"
	"github.com/hobord/poc-htmx-go-todolist/dal/config/viper"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web"
	"github.com/hobord/poc-htmx-go-todolist/services/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer cancel()

	v := viper.NewReader()
	config := config.NewService(v)

	conf, err := config.GetServerConfig()
	if err != nil {
		os.Exit(1)
	}

	services, err := composition.NewServerServices(ctx, conf)
	if err != nil {
		os.Exit(1)
	}

	log := services.Log

	webServer, err := web.NewServer(ctx, conf, services)
	if err != nil {
		log.Error("Could not create web server", "error", err)
		os.Exit(1)
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return webServer.Start(gCtx)
	})

	g.Go(func() error {
		<-gCtx.Done()

		if err := webServer.Stop(context.Background()); err != nil {
			log.Error("Could not stop web server", "error", err)
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error("Shutting down server", "error", err)
	}
}
