package main

import (
	"log"

	"github.com/mauricewittek/go-grpc-service/internal/db"
	"github.com/mauricewittek/go-grpc-service/internal/rocket"
	"github.com/mauricewittek/go-grpc-service/internal/transport/grpc"
)

func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
