package broker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi/v11/domain"
	"golang.org/x/net/context"

	"github.com/cloudfoundry/cloud-service-broker/v2/brokerapi/broker/brokerfakes"
	"github.com/cloudfoundry/cloud-service-broker/v2/utils"

	"github.com/cloudfoundry/cloud-service-broker/v2/brokerapi/broker"
)

var _ = Describe("GetInstance", func() {
	It("is not implemented", func() {
		serviceBroker, err := broker.New(&broker.BrokerConfig{}, &brokerfakes.FakeStorage{}, utils.NewLogger("brokers-test"))
		Expect(err).ToNot(HaveOccurred())

		_, err = serviceBroker.GetInstance(context.TODO(), "instance-id", domain.FetchInstanceDetails{})

		Expect(err).To(MatchError("the service_instances endpoint is unsupported"))
	})
})
