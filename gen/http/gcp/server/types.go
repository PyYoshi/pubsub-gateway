// Code generated by goa v3.7.5, DO NOT EDIT.
//
// gcp HTTP server types
//
// Command:
// $ goa gen github.com/PyYoshi/pubsub-gateway/design -o ./

package server

import (
	"unicode/utf8"

	gcp "github.com/PyYoshi/pubsub-gateway/gen/gcp"
	gcpviews "github.com/PyYoshi/pubsub-gateway/gen/gcp/views"
	goa "goa.design/goa/v3/pkg"
)

// PublishRequestBody is the type of the "gcp" service "publish" endpoint HTTP
// request body.
type PublishRequestBody struct {
	// Cloud Pub/SubのメッセージBody
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// PublishResponseBody is the type of the "gcp" service "publish" endpoint HTTP
// response body.
type PublishResponseBody struct {
	// message id
	ID string `form:"id" json:"id" xml:"id"`
}

// NewPublishResponseBody builds the HTTP response body from the result of the
// "publish" endpoint of the "gcp" service.
func NewPublishResponseBody(res *gcpviews.PubsubGatewayGcpPublishView) *PublishResponseBody {
	body := &PublishResponseBody{
		ID: *res.ID,
	}
	return body
}

// NewPublishGcpPublishRequestType builds a gcp service publish endpoint
// payload.
func NewPublishGcpPublishRequestType(body *PublishRequestBody, topic string) *gcp.GcpPublishRequestType {
	v := &gcp.GcpPublishRequestType{
		Message: *body.Message,
	}
	v.Topic = topic

	return v
}

// ValidatePublishRequestBody runs the validations defined on PublishRequestBody
func ValidatePublishRequestBody(body *PublishRequestBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Message != nil {
		if utf8.RuneCountInString(*body.Message) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.message", *body.Message, utf8.RuneCountInString(*body.Message), 1, true))
		}
	}
	return
}
