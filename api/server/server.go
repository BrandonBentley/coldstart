package server

import (
	"context"
	"fmt"

	"github.com/BrandonBentley/coldstart/api/handler/health"
	"github.com/BrandonBentley/coldstart/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In
	fx.Lifecycle

	Config *config.Config

	// Exported Handlers
	HealthHandler *health.Handler
}

type Server struct {
	engine *gin.Engine

	// Unexported Handlers
	healthHandler *health.Handler

	config *config.Config
}

func NewServer(params ServerParams) *Server {
	s := &Server{
		engine:        gin.Default(),
		healthHandler: params.HealthHandler,
		config:        params.Config,
	}

	s.registerEndpoints()

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.engine.Run(fmt.Sprintf(":%d", s.config.Server.Http.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return s
}

func (s *Server) registerEndpoints() {
	s.engine.GET("/health", s.healthHandler.Check)
}
