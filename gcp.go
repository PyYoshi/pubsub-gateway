package pubusubgateway

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/PyYoshi/cloud_pubsub_gateway/config"
	gcp "github.com/PyYoshi/cloud_pubsub_gateway/gen/gcp"
	log "github.com/PyYoshi/cloud_pubsub_gateway/gen/log"
)

// gcp service example implementation.
// The example methods log the requests and return zero values.
type gcpsrvc struct {
	logger    *log.Logger
	cfg       *config.GatewayServer
	pubsubCli *pubsub.Client
}

// NewGcp returns the gcp service implementation.
func NewGcp(logger *log.Logger, cfg *config.GatewayServer, pubsubCli *pubsub.Client) gcp.Service {
	return &gcpsrvc{
		logger:    logger,
		cfg:       cfg,
		pubsubCli: pubsubCli,
	}
}

// httpプロトコルを経由してCloud Pub/Subにメッセージを渡すエンドポイント
func (s *gcpsrvc) Publish(ctx context.Context, p *gcp.GcpPublishRequestType) (res *gcp.PubsubGatewayGcpPublish, err error) {
	_l := s.logger.Named("gcpsrvc.Publish")

	topic := s.pubsubCli.Topic(p.Topic)
	defer topic.Stop()

	pres := topic.Publish(
		ctx,
		&pubsub.Message{
			Data: []byte(p.Message),
		},
	)
	messageID, err := pres.Get(ctx)
	if err != nil {
		return nil, err
	}

	_l.Infow(
		"Successfully published a message",
		"topic",
		p.Topic,
		"message",
		p.Message,
		"message_id",
		messageID,
	)

	res = &gcp.PubsubGatewayGcpPublish{
		ID: messageID,
	}
	return
}
