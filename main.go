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

	sub := subscriber.New(client)
	_, err = sub.Subscribe(ctx, subscriber.Options{TopicName: "project-id"})
	if err != nil {
		log.Fatal(err)
	}
}
