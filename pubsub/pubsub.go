package pubsub

import (
	"context"
	"time"

	gpubsub "cloud.google.com/go/pubsub"
)

type Client interface {
	CreateTopic(ctx context.Context, topicID string) (Topic, error)
	Topic(id string) Topic
	CreateSubscription(ctx context.Context, id string, cfg SubscriptionConfig) (Subscription, error)
	Subscription(id string) Subscription
}

type Topic interface {
	String() string
}

type Subscription interface {
	Exists(ctx context.Context) (bool, error)
	Receive(ctx context.Context, f func(context.Context, Message)) error
	Delete(ctx context.Context) error
}

type SubscriptionConfig struct {
	gpubsub.SubscriptionConfig
	Topic Topic
}

type Message interface {
	ID() string
	Data() []byte
	Attributes() map[string]string
	PublishTime() time.Time
	Ack()
	Nack()
}

type PubSub struct {
	client Client
}

type Subscriber struct {
	subscription Subscription
}

type Options struct {
	ProjectID        string
	TopicName        string
	SubscriptionName string
}

func (ps *PubSub) Subscribe(ctx context.Context, opts Options) (*Subscriber, error) {
	topic, err := ps.client.CreateTopic(ctx, opts.TopicName)
	if err != nil {
		return &Subscriber{}, err
	}

	sub := ps.client.Subscription(opts.SubscriptionName)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return &Subscriber{}, err
	}

	if !exists {
		sub, err = ps.client.CreateSubscription(ctx, opts.SubscriptionName, SubscriptionConfig{Topic: topic})
		if err != nil {
			return &Subscriber{}, err
		}
	}

	return &Subscriber{subscription: sub}, nil
}
