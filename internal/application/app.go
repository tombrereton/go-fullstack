package application

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	port := "4000"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	fmt.Printf("Starting server on http://localhost:%s\n", port)
	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("error starting server: %w", err)
		}
		close(ch)
	}()

	select {
	case err :=  <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
