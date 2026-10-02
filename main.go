package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Ficserbiyy/weather-api/internal/application"
)

func main() {
	app := application.NewApp(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	cancel()

	log.Println("Server listening on 0.0.0.0:8080")
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
