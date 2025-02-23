package client

import (
	"github.com/BrandonBentley/coldstart/client/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"client",
	fx.Provide(
		health.NewClient,
	),
)
