package serve

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"../../config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/martian/log"
)

func Run(cfg *config.Config) {
	// initiate chie as router agent
	r := chi.NewRouter()
	// set timeout
	r.Use(middleware.Timeout(60 * time.Second))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Errorf("app - Run - httpServer.Notify: %w", err)
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Infof("app - Run - signal: %s", s.String())
	default:
	}
}
