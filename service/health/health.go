package health

import (
	"github.com/BrandonBentley/coldstart/client/health"
	"github.com/BrandonBentley/coldstart/client/health/healthintf"
	healthe "github.com/BrandonBentley/coldstart/entity/health"
	healtheintf "github.com/BrandonBentley/coldstart/entity/health/healthintf"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	HealthClient *health.Client
	HealthEntity *healthe.Entity
}

type Service struct {
	healthClient healthintf.Client
	healthEntity healtheintf.Entity
}

func NewService(params ServiceParams) *Service {
	return &Service{
		healthClient: params.HealthClient,
		healthEntity: params.HealthEntity,
	}
}
