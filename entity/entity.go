package entity

import (
	"github.com/BrandonBentley/coldstart/entity/health"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"entity",
	fx.Provide(
		health.NewEntity,
	),
)
