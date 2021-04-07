package health

import (
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/health"
)

type healthService struct {
	pb.UnimplementedHealthServer // embedding
}

// NewHealthService generates the `HealthServer` implementation.
func NewHealthService() pb.HealthServer {
	return &healthService{}
}
