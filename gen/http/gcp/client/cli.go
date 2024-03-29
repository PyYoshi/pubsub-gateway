// Code generated by goa v3.13.1, DO NOT EDIT.
//
// gcp HTTP client CLI support package
//
// Command:
// $ goa gen github.com/PyYoshi/pubsub-gateway/design -o ./

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	gcp "github.com/PyYoshi/pubsub-gateway/gen/gcp"
	goa "goa.design/goa/v3/pkg"
)

// BuildPublishPayload builds the payload for the gcp publish endpoint from CLI
// flags.
func BuildPublishPayload(gcpPublishBody string, gcpPublishTopic string) (*gcp.GcpPublishRequestType, error) {
	var err error
	var body PublishRequestBody
	{
		err = json.Unmarshal([]byte(gcpPublishBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"message\": \"166\"\n   }'")
		}
		if utf8.RuneCountInString(body.Message) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.message", body.Message, utf8.RuneCountInString(body.Message), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var topic string
	{
		topic = gcpPublishTopic
		if utf8.RuneCountInString(topic) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("topic", topic, utf8.RuneCountInString(topic), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &gcp.GcpPublishRequestType{
		Message: body.Message,
	}
	v.Topic = topic

	return v, nil
}
