package main

import (
	stan_sub "first_task_l0/internal/clients/stan-sub"
	"first_task_l0/internal/config"
	"first_task_l0/internal/events"
	"first_task_l0/internal/storage/postgres"
	"fmt"
	"log/slog"
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
		log.Error(fmt.Sprintf("can't subscribe to topic %s:", subject), err)
	}
	storage, err := postgres.New(cfg)
	if err != nil {
		log.Error("can't init postgres storage")
		os.Exit(1)
	}

	processor := events.New(storage)

	for order := range subscriber.Orders {
		log.Info(fmt.Sprintf("get order %s from channel", order.Name))
		err = processor.CreateOrder(order)
		if err != nil {
			log.Error("can't create order", err)
		} else {
			log.Info("create order")
		}
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
