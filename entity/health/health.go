package health

import "go.uber.org/fx"

type EntityParams struct {
	fx.In
}

type Entity struct {
	// DB
}

func NewEntity() *Entity {
	return &Entity{}
}
