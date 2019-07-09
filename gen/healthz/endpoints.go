// Code generated by goa v3.0.2, DO NOT EDIT.
//
// healthz endpoints
//
// Command:
// $ goa gen github.com/PyYoshi/cloud_pubsub_gateway/design -o ./

package healthz

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "healthz" service endpoints.
type Endpoints struct {
	Readiness goa.Endpoint
	Liveness  goa.Endpoint
}

// NewEndpoints wraps the methods of the "healthz" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Readiness: NewReadinessEndpoint(s),
		Liveness:  NewLivenessEndpoint(s),
	}
}

// Use applies the given middleware to all the "healthz" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Readiness = m(e.Readiness)
	e.Liveness = m(e.Liveness)
}

// NewReadinessEndpoint returns an endpoint function that calls the method
// "readiness" of service "healthz".
func NewReadinessEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Readiness(ctx)
	}
}

// NewLivenessEndpoint returns an endpoint function that calls the method
// "liveness" of service "healthz".
func NewLivenessEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Liveness(ctx)
	}
}