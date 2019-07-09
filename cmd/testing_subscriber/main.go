package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

type config struct {
	gcpProjectID       string
	pubsubTopic        string
	pubsubSubscription string
}

func getConfig() *config {
	cfg := &config{}
	cfg.gcpProjectID = os.Getenv("GATEWAY_SERVER_GOOGLE_PROJECT_ID")
	cfg.pubsubTopic = os.Getenv("PUBSUB_TOPIC")
	cfg.pubsubSubscription = os.Getenv("PUBSUB_SUBSCRIPTION")
	return cfg
}

func main() {
	cfg := getConfig()

	ctx := context.Background()
	pubsubCli, err := pubsub.NewClient(ctx, cfg.gcpProjectID)
	if err != nil {
		panic(err)
	}
	defer pubsubCli.Close()

	topic := pubsubCli.Topic(cfg.pubsubTopic)
	topicExists, err := topic.Exists(ctx)
	if err != nil {
		panic(err)
	}
	if !topicExists {
		topic, err = pubsubCli.CreateTopic(ctx, cfg.pubsubTopic)
		if err != nil {
			panic(err)
		}
	}
	log.Printf("topic: %#v\n", topic)

	subsc := pubsubCli.Subscription(cfg.pubsubSubscription)
	subscExists, err := subsc.Exists(ctx)
	if err != nil {
		panic(err)
	}
	if !subscExists {
		subsc, err = pubsubCli.CreateSubscription(ctx, cfg.pubsubSubscription, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			panic(err)
		}
	}
	log.Printf("subscription: %#v\n", subsc)

	err = subsc.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("message data: %#v\n", string(msg.Data))
	})
	if err != nil {
		panic(err)
	}
}
