package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	dev   = "dev"
	stage = "stage"
	prod  = "prod"
)

const (
	defaultConfigBasePath = "config/env"
)

func NewConfig() (*Config, error) {
	cfg := viper.New()
	cfg.SetConfigType("json")
	cfg.SetConfigFile(getConfigPath())

	bindEnvs(cfg)
	setDefaults(cfg)

	err := cfg.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = cfg.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigPath() string {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case dev, stage, prod:
	default:
		appEnv = dev
	}
	return path.Join(getConfigBasePath(), appEnv+".json")
}

func getConfigBasePath() string {
	basePath := os.Getenv("CONFIG_BASE_PATH")
	if basePath == "" {
		return defaultConfigBasePath
	}
	return basePath
}
