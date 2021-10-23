package serve

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"../../config"
)

func Execute() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	Run(&cfg)
}
