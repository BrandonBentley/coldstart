package main

import (
	"github.com/BrandonBentley/coldstart/api"
	"github.com/BrandonBentley/coldstart/client"
	"github.com/BrandonBentley/coldstart/config"
	"github.com/BrandonBentley/coldstart/entity"
	"github.com/BrandonBentley/coldstart/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
		),
		api.Module,
		client.Module,
		entity.Module,
		service.Module,
	)

	app.Run()
}
