package main

import (
	"log"

	"github.com/mauricewittek/go-grpc-service/internal/db"
	"github.com/mauricewittek/go-grpc-service/internal/rocket"
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

	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
