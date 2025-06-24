package main

import (
	"context"
	"github.com/nk-31012002/student-api/internal/config"
	"github.com/nk-31012002/student-api/internal/http/handlers/students"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//setup database
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", students.New())
	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started %s", slog.String("Address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("shut down successfully")
}
