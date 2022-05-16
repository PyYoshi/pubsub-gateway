// Code generated by goa v3.7.5, DO NOT EDIT.
//
// gcp service
//
// Command:
// $ goa gen github.com/PyYoshi/pubsub-gateway/design -o ./

package gcp

import (
	"context"

	gcpviews "github.com/PyYoshi/pubsub-gateway/gen/gcp/views"
)

// Google Cloud Pub/Sub向けサービス
type Service interface {
	// httpプロトコルを経由してCloud Pub/Subにメッセージを渡すエンドポイント
	Publish(context.Context, *GcpPublishRequestType) (res *PubsubGatewayGcpPublish, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "gcp"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"publish"}

// GcpPublishRequestType is the payload type of the gcp service publish method.
type GcpPublishRequestType struct {
	// Cloud Pub/SubのTopic
	Topic string
	// Cloud Pub/SubのメッセージBody
	Message string
}

// PubsubGatewayGcpPublish is the result type of the gcp service publish method.
type PubsubGatewayGcpPublish struct {
	// message id
	ID string
}

// NewPubsubGatewayGcpPublish initializes result type PubsubGatewayGcpPublish
// from viewed result type PubsubGatewayGcpPublish.
func NewPubsubGatewayGcpPublish(vres *gcpviews.PubsubGatewayGcpPublish) *PubsubGatewayGcpPublish {
	return newPubsubGatewayGcpPublish(vres.Projected)
}

// NewViewedPubsubGatewayGcpPublish initializes viewed result type
// PubsubGatewayGcpPublish from result type PubsubGatewayGcpPublish using the
// given view.
func NewViewedPubsubGatewayGcpPublish(res *PubsubGatewayGcpPublish, view string) *gcpviews.PubsubGatewayGcpPublish {
	p := newPubsubGatewayGcpPublishView(res)
	return &gcpviews.PubsubGatewayGcpPublish{Projected: p, View: "default"}
}

// newPubsubGatewayGcpPublish converts projected type PubsubGatewayGcpPublish
// to service type PubsubGatewayGcpPublish.
func newPubsubGatewayGcpPublish(vres *gcpviews.PubsubGatewayGcpPublishView) *PubsubGatewayGcpPublish {
	res := &PubsubGatewayGcpPublish{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	return res
}

// newPubsubGatewayGcpPublishView projects result type PubsubGatewayGcpPublish
// to projected type PubsubGatewayGcpPublishView using the "default" view.
func newPubsubGatewayGcpPublishView(res *PubsubGatewayGcpPublish) *gcpviews.PubsubGatewayGcpPublishView {
	vres := &gcpviews.PubsubGatewayGcpPublishView{
		ID: &res.ID,
	}
	return vres
}
