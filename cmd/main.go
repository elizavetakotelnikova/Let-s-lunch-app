package main

import (
	"cmd/app/config"
	"cmd/di"
	"context"
	"log"
)

func main() {
	container, err := di.NewContainer(
		config.Params{DatabaseURL: "postgresql://localhost/postgres?user=admin&password=admin"},
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := container.Server(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	container.Close()
}
