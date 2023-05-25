package subscriber

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
)

type Client interface {
	CreateTopic(ctx context.Context, topicID string) (Topic, error)
	Topic(id string) Topic
	CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (Subscription, error)
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

func New(client Client) *PubSub {
	return &PubSub{client: client}
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
		sub, err = ps.client.CreateSubscription(ctx, opts.SubscriptionName, pubsub.SubscriptionConfig{Topic: topic.(*pubsub.Topic)})
		if err != nil {
			return &Subscriber{}, err
		}
	}

	return &Subscriber{subscription: sub}, nil
}
