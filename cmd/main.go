package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/tombrereton/go-hot-reload/application"
)

func main() {
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err == nil {
		fmt.Println("Shutting down server: %w", err)
	}
}
