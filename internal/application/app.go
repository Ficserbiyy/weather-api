package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router *gin.Engine
	rdb    *redis.Client
	cfg    config
}

func NewApp(config config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		cfg: config,
	}

	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("unable to ping redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			log.Println("unable to close redis: ", err)
		}
	}()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.ServerPort),
		Handler: a.router,
	}
	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			ch <- fmt.Errorf("unable to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err

	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
