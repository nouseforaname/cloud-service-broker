package testdrive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pivotal-cf/brokerapi/v11/domain"
)

func (b *Broker) LastOperation(serviceInstanceGUID string) (domain.LastOperation, error) {
	lastOperationResponse := b.Client.LastOperation(serviceInstanceGUID, uuid.NewString())
	switch {
	case lastOperationResponse.Error != nil:
		return domain.LastOperation{}, lastOperationResponse.Error
	case lastOperationResponse.StatusCode != http.StatusOK:
		return domain.LastOperation{}, &UnexpectedStatusError{StatusCode: lastOperationResponse.StatusCode, ResponseBody: lastOperationResponse.ResponseBody}
	}

	var receiver domain.LastOperation
	if err := json.Unmarshal(lastOperationResponse.ResponseBody, &receiver); err != nil {
		return domain.LastOperation{}, err
	}

	return receiver, nil
}

func (b *Broker) LastOperationFinalState(serviceInstanceGUID string) (domain.LastOperationState, error) {
	lastOperation, err := b.LastOperationFinalValue(serviceInstanceGUID)
	if err != nil {
		return "", err
	}
	return lastOperation.State, nil
}

func (b *Broker) LastOperationFinalValue(serviceInstanceGUID string) (domain.LastOperation, error) {
	start := time.Now()
	for {
		lastOperation, err := b.LastOperation(serviceInstanceGUID)
		switch {
		case err != nil:
			return domain.LastOperation{}, err
		case time.Since(start) > time.Hour:
			return domain.LastOperation{}, fmt.Errorf("timed out waiting for last operation on service instance %q", serviceInstanceGUID)
		case lastOperation.State == domain.Failed, lastOperation.State == domain.Succeeded:
			return lastOperation, nil
		default:
			time.Sleep(time.Second)
		}
	}
}

func toJSONRawMessage(params any) (json.RawMessage, error) {
	switch p := params.(type) {
	case nil:
		return nil, nil
	case string:
		return json.RawMessage(p), nil
	case []byte:
		return p, nil
	default:
		return json.Marshal(p)
	}
}
