package http

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/bariasabda/monorepo/packages/fetch/config"
	"github.com/bariasabda/monorepo/packages/fetch/domain/repository"
	"github.com/bariasabda/monorepo/packages/fetch/domain/service"
)

func Execute() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./packages/fetch/config/config.yml", &cfg)
	if err != nil {
		err = cleanenv.ReadConfig("./config/config.yml", &cfg)
		if err != nil {
			log.Fatalf("Config error: %s", err)
		}
	}

	repository := repository.NewRepository(cfg)
	service := service.NewService(cfg, repository)

	// Run
	handler := NewHandler(service)
	r := NewRouter(*handler)
	r.routes()
	r.router.Run(cfg.Base.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatalf("app - Run - signal: %s", s.String())
	default:
	}
}
