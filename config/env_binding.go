package config

import "github.com/spf13/viper"

func setDefaults(cfg *viper.Viper) {
	cfg.SetDefault("app.env", dev)
}

func bindEnvs(cfg *viper.Viper) {
	cfg.BindEnv("app.env", "APP_ENV")
}
