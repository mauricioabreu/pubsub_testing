package subscriber

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type PubSubClient interface {
	CreateTopic(ctx context.Context, topic string) (*pubsub.Topic, error)
	Subscription(name string) *pubsub.Subscription
	CreateSubscription(ctx context.Context, name string, config pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
}

type PubSub struct {
	client PubSubClient
}

func New(client PubSubClient) *PubSub {
	return &PubSub{client: client}
}

type Options struct {
	ProjectID        string
	TopicName        string
	SubscriptionName string
}

func (ps *PubSub) Subscribe(ctx context.Context, opts Options) (*pubsub.Subscription, error) {
	topic, err := ps.client.CreateTopic(ctx, opts.TopicName)
	if err != nil {
		return &pubsub.Subscription{}, err
	}

	sub := ps.client.Subscription(opts.SubscriptionName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return &pubsub.Subscription{}, err
	}

	if !exists {
		sub, err = ps.client.CreateSubscription(ctx, opts.SubscriptionName, pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			return &pubsub.Subscription{}, err
		}
	}

	return sub, nil
}
