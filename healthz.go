package pubusubgateway

import (
	"context"

	healthz "github.com/PyYoshi/cloud_pubsub_gateway/gen/healthz"
	log "github.com/PyYoshi/cloud_pubsub_gateway/gen/log"
)

// healthz service example implementation.
// The example methods log the requests and return zero values.
type healthzsrvc struct {
	logger *log.Logger
}

// NewHealthz returns the healthz service implementation.
func NewHealthz(logger *log.Logger) healthz.Service {
	return &healthzsrvc{logger}
}

// サーバが準備完了状態でいるかチェックするエンドポイント
func (s *healthzsrvc) Readiness(ctx context.Context) (err error) {
	s.logger.Named("healthzsrvc.Readiness").Info("nothing to do")
	return nil
}

// サーバが生きているかチェックするエンドポイント
func (s *healthzsrvc) Liveness(ctx context.Context) (err error) {
	s.logger.Named("healthzsrvc.Liveness").Info("nothing to do")
	return nil
}
