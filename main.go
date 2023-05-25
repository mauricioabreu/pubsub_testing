package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/mauricioabreu/pubsub_testing/subscriber"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		log.Fatal(err)
	}

	adapter := &subscriber.PubSubClientAdapter{Client: client}
	sub := subscriber.New(adapter)
	_, err = sub.Subscribe(ctx, subscriber.Options{TopicName: "topic-id"})
	if err != nil {
		log.Fatal(err)
	}
}
