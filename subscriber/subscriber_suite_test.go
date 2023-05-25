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
		ctrl    *gomock.Controller
		mclient *mock_subscriber.MockClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mclient = mock_subscriber.NewMockClient(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Creating topic", func() {
		It("Creates a topic", func() {
			ctx := context.TODO()
			mclient.EXPECT().CreateTopic(ctx, "project-id").Return(&pubsub.Topic{}, errors.New("failed to create topic"))
			psub := subscriber.New(mclient)
			_, err := psub.Subscribe(ctx, subscriber.Options{TopicName: "project-id"})
			Expect(err).To(HaveOccurred())
		})
	})
})
