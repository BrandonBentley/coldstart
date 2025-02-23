package service

import (
	"github.com/BrandonBentley/coldstart/service/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Provide(
		health.NewService,
	),
)
