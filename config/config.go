package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(
		NewConfig,
	),
)

const (
	AppDev     = "dev"
	AppStaging = "staging"
	AppProd    = "prod"
)

func NewConfig() (*Config, error) {
	loadDotEnv()

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg.Server.Http.Port)

	return &cfg, nil
}

func loadDotEnv() {
	switch os.Getenv("APP_ENV") {
	case "", AppDev:
		godotenv.Load("config/env/dev.env")
	}
}
