package subscriber_test

import (
	"errors"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/golang/mock/gomock"
	mock_subscriber "github.com/mauricioabreu/pubsub_testing/mocks"
	"github.com/mauricioabreu/pubsub_testing/subscriber"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

func TestPubsub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pubsub Suite")
}

var _ = Describe("Subscription service", func() {
	var (
		ctrl          *gomock.Controller
		mclient       *mock_subscriber.MockPubSubClient
		msubscription *mock_subscriber.MockSubscription
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mclient = mock_subscriber.NewMockPubSubClient(ctrl)
		msubscription = mock_subscriber.NewMockSubscription(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Subscribing to a topic", func() {
		It("Fails creating a topic", func() {
			ctx := context.TODO()
			mclient.EXPECT().CreateTopic(ctx, "topic-id").Return(&pubsub.Topic{}, errors.New("failed to create topic"))
			psub := subscriber.New(mclient)
			_, err := psub.Subscribe(ctx, subscriber.Options{TopicName: "topic-id"})
			Expect(err).To(HaveOccurred())
		})
		It("Creates a topic but fetching subscription fails", func() {
			ctx := context.TODO()
			mclient.EXPECT().CreateTopic(ctx, "topic-id").Return(&pubsub.Topic{}, nil)
			msubscription.EXPECT().Exists(ctx).Return(false, errors.New("failed to fetch subscription"))
			mclient.EXPECT().Subscription("subscription-id").Return(msubscription)
			psub := subscriber.New(mclient)
			_, err := psub.Subscribe(ctx, subscriber.Options{TopicName: "topic-id", SubscriptionName: "subscription-id"})
			Expect(err).To(HaveOccurred())
		})
	})
})
