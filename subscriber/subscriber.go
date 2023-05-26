package subscriber

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// PubSubClient interface abstracts the pubsub.Client dependency
type PubSubClient interface {
	CreateTopic(ctx context.Context, topic string) (*pubsub.Topic, error)
	Subscription(name string) Subscription
	CreateSubscription(ctx context.Context, name string, config pubsub.SubscriptionConfig) (Subscription, error)
}

// Subscription interface abstracts the pubsub.Subscription dependency
type Subscription interface {
	Exists(ctx context.Context) (bool, error)
}

// PubSubClientAdapter is an adapter that wraps the pubsub.Client and implements the PubSubClient interface
type PubSubClientAdapter struct {
	Client *pubsub.Client
}

func (a *PubSubClientAdapter) CreateTopic(ctx context.Context, topic string) (*pubsub.Topic, error) {
	return a.Client.CreateTopic(ctx, topic)
}

func (a *PubSubClientAdapter) Subscription(name string) Subscription {
	return a.Client.Subscription(name)
}

func (a *PubSubClientAdapter) CreateSubscription(ctx context.Context, name string, config pubsub.SubscriptionConfig) (Subscription, error) {
	return a.Client.CreateSubscription(ctx, name, config)
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

func (ps *PubSub) Subscribe(ctx context.Context, opts Options) (Subscription, error) {
	topic, err := ps.client.CreateTopic(ctx, opts.TopicName)
	if err != nil {
		return nil, err
	}

	sub := ps.client.Subscription(opts.SubscriptionName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !exists {
		sub, err = ps.client.CreateSubscription(ctx, opts.SubscriptionName, pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			return nil, err
		}
	}

	return sub, nil
}
