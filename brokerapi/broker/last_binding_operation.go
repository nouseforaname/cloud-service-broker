package broker

import (
	"context"

	"code.cloudfoundry.org/lager/v3"
	"github.com/pivotal-cf/brokerapi/v11/domain"
	"github.com/pivotal-cf/brokerapi/v11/domain/apiresponses"

	"github.com/cloudfoundry/cloud-service-broker/v2/utils/correlation"
)

// LastBindingOperation fetches last operation state for a service binding.
// GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation
//
// NOTE: This functionality is not implemented.
func (broker *ServiceBroker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details domain.PollDetails) (domain.LastOperation, error) {
	broker.Logger.Info("LastBindingOperation", correlation.ID(ctx), lager.Data{
		"instance_id":    instanceID,
		"binding_id":     bindingID,
		"plan_id":        details.PlanID,
		"service_id":     details.ServiceID,
		"operation_data": details.OperationData,
	})

	return domain.LastOperation{}, apiresponses.ErrAsyncRequired
}
