// Code generated by goa v3.0.2, DO NOT EDIT.
//
// healthz client
//
// Command:
// $ goa gen github.com/PyYoshi/cloud_pubsub_gateway/design -o ./

package healthz

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "healthz" service client.
type Client struct {
	ReadinessEndpoint goa.Endpoint
	LivenessEndpoint  goa.Endpoint
}

// NewClient initializes a "healthz" service client given the endpoints.
func NewClient(readiness, liveness goa.Endpoint) *Client {
	return &Client{
		ReadinessEndpoint: readiness,
		LivenessEndpoint:  liveness,
	}
}

// Readiness calls the "readiness" endpoint of the "healthz" service.
func (c *Client) Readiness(ctx context.Context) (err error) {
	_, err = c.ReadinessEndpoint(ctx, nil)
	return
}

// Liveness calls the "liveness" endpoint of the "healthz" service.
func (c *Client) Liveness(ctx context.Context) (err error) {
	_, err = c.LivenessEndpoint(ctx, nil)
	return
}
