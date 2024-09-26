package main

import (
	"context"
	"log"

	"github.com/kyrare/ya-diplom-2/internal/app/client"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
)

func main() {
	ctx := context.Background()

	config, err := services.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	l, err := services.NewLogger(config.Debug)
	if err != nil {
		log.Fatal(err)
	}

	app := client.NewApp(config, l)

	if err := app.Configure(ctx); err != nil {
		l.Fatal(err)
		return
	}

	if err := app.Run(ctx); err != nil {
		l.Fatal(err)
		return
	}
}
