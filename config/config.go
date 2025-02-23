package config

type Config struct {
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
}

type App struct {
	Env string `mapstructure:"env"`
}

type Server struct {
	Http Http `mapstructure:"http"`
}

type Http struct {
	Port int `mapstructure:"port"`
}
