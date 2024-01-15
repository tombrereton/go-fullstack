package application

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	cfg    Configuration
	router http.Handler
}

func NewApp(c Configuration) *App {
	app := &App{
		router: loadRoutes(),
		cfg:    c,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	port := a.cfg.port
	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("error starting server: %w", err)
		}
		close(ch)
	}()

	fmt.Printf("Started server at http://localhost:%s\n", port)
	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
