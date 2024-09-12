package main

import (
	"context"
	"log"

	"github.com/kyrare/ya-diplom-2/internal/app/server"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
)

func main() {
	ctx := context.Background()

	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	l, err := services.NewLogger(false)
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(config, l)

	if err := app.Configure(ctx); err != nil {
		l.Fatal(err)
		return
	}

	if err := app.Run(ctx); err != nil {
		l.Fatal(err)
		return
	}
}
