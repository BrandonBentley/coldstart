package main

import (
	"github.com/BrandonBentley/coldstart/api"
	"github.com/BrandonBentley/coldstart/client"
	"github.com/BrandonBentley/coldstart/config"
	"github.com/BrandonBentley/coldstart/entity"
	"github.com/BrandonBentley/coldstart/service"
	_ "github.com/BrandonBentley/slogctx/sloginit"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		api.Module,
		client.Module,
		config.Module,
		entity.Module,
		service.Module,
	)
	app.Run()
}
