package healthintf

import (
	"context"

	"github.com/BrandonBentley/coldstart/service/health"
)

//go:generate mockgen -package=health -source=health.go -destination=../mock_health/health.go
type Service interface {
	Check(ctx context.Context) (*health.HealthCheckResponse, error)
}
