package main

import (
	stan_sub "first_task_l0/internal/clients/stan-sub"
	"first_task_l0/internal/config"
	"first_task_l0/internal/events"
	"first_task_l0/internal/http_server/handlers/order"
	"first_task_l0/internal/lib"
	cache2 "first_task_l0/internal/storage/cache"
	"first_task_l0/internal/storage/postgres"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

const (
	clusterID = "test-cluster"
	clientID  = "my_client"
	subject   = "orders"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting demo-service", slog.String("env", cfg.Env))

	subscriber := stan_sub.New(clusterID, clientID, log)
	err := subscriber.Subscribe(subject)
	if err != nil {
		log.Error(fmt.Sprintf("can't subscribe to topic %s:", subject), lib.Err(err))
	}

	storage, err := postgres.New(cfg)
	if err != nil {
		log.Error("can't init postgres storage")
		os.Exit(1)
	}

	cache, err := cache2.New(cfg)
	if err != nil {
		log.Error("can't init cache", lib.Err(err))
		os.Exit(1)
	}

	err = cache.Recovery()
	if err != nil {
		log.Error("can't recovery cache", lib.Err(err))
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/orders", func(r chi.Router) {
		r.Get("/{orderID}", order.New(log, cache))
	})

	go http.ListenAndServe(":3333", r)

	processor := events.New(storage)

	for order := range subscriber.Orders {
		err = processor.CreateOrder(order)
		if err != nil {
			log.Error("can't create order %v", lib.Err(err))
		} else {
			log.Info(fmt.Sprintf("create order %s in db", order.OrderUID))
		}

		cache.CreateOrder(order)
		log.Info(fmt.Sprintf("create order %s in cache", order.OrderUID))
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
