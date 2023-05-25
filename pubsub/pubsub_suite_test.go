package pubsub_test

import (
	"errors"
	"testing"

	gpubsub "cloud.google.com/go/pubsub"
	"github.com/golang/mock/gomock"
	mock_pubsub "github.com/mauricioabreu/pubsub_testing/mocks"
	"github.com/mauricioabreu/pubsub_testing/pubsub"
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
		mclient *mock_pubsub.MockClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mclient = mock_pubsub.NewMockClient(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Creating topic", func() {
		It("Creates a topic", func() {
			ctx := context.TODO()
			mclient.EXPECT().CreateTopic(ctx, "project-id").Return(&gpubsub.Topic{}, errors.New("failed to create topic"))
			psub := pubsub.New(mclient)
			_, err := psub.Subscribe(ctx, pubsub.Options{TopicName: "project-id"})
			Expect(err).To(HaveOccurred())
		})
	})
})
