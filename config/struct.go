package config

type Config struct {
	App    App    `envPrefix:"APP_"`
	Server Server `envPrefix:"SERVER_"`
}

type App struct {
	Env string `env:"ENV" envDefault:"dev"`
}

type Server struct {
	Http Http `envPrefix:"HTTP_"`
}

type Http struct {
	Port int `env:"PORT" envDefault:"80"`
}
