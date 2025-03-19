package config

import (
	"os"
	"path"

	"github.com/BrandonBentley/ezbind"
	"github.com/spf13/viper"
)

const (
	AppDev     = "dev"
	AppStaging = "staging"
	AppProd    = "prod"
)

const (
	defaultConfigBasePath = "config/env"
)

func NewConfig() (*Config, error) {
	cfg := viper.New()
	cfg.SetConfigType("json")
	cfg.SetConfigFile(getConfigPath())

	setDefaults(cfg)

	err := cfg.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	ezbind.BindStruct(cfg, config)

	err = cfg.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigPath() string {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case AppDev, AppStaging, AppProd:
	default:
		appEnv = AppDev
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
