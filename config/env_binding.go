package config

import "github.com/spf13/viper"

func setDefaults(cfg *viper.Viper) {
	cfg.SetDefault("app.env", AppDev)
	cfg.SetDefault("server.http.port", 80)
}
