package api

import (
	"github.com/BrandonBentley/coldstart/api/handler"
	"github.com/BrandonBentley/coldstart/api/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	fx.Invoke(
		server.NewServer,
	),
	handler.Module,
)
