package healthintf

import "context"

//go:generate mockgen -package=mock_health -source=health.go -destination=../mock_health/health.go
type Client interface {
	Ping(ctx context.Context) error
}
