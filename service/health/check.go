package health

import "context"

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func (s *Service) Check(ctx context.Context) (*HealthCheckResponse, error) {
	err := s.healthClient.Ping(ctx)
	if err != nil {
		return nil, err
	}
	err = s.healthEntity.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &HealthCheckResponse{
		Status: "ok",
	}, nil
}
