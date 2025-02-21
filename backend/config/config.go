package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	OpenAIKey   string `env:"OPEN_AI_KEY"`
	DatabaseURL string `env:"DATABASE_URL"`
	Port        int    `env:"PORT" envDefault:"8080"`
}

var Config config

func Init() {
	godotenv.Load()

	if err := env.ParseWithOptions(&Config, env.Options{
		RequiredIfNoDef: true,
	}); err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}
}
