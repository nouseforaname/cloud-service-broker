// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package brokers

import (
	"fmt"

	"code.cloudfoundry.org/lager"

	"github.com/cloudfoundry-incubator/cloud-service-broker/pkg/broker"
	"github.com/cloudfoundry-incubator/cloud-service-broker/pkg/brokerpak"
	"github.com/cloudfoundry-incubator/cloud-service-broker/pkg/config"
	"github.com/cloudfoundry-incubator/cloud-service-broker/pkg/credstore"
)

type BrokerConfig struct {
	Registry   broker.BrokerRegistry
	Credstore  credstore.CredStore
}

func NewBrokerConfigFromEnv(logger lager.Logger) (*BrokerConfig, error) {
	registry := broker.BrokerRegistry{}
	if err := brokerpak.RegisterAll(registry); err != nil {
		return nil, fmt.Errorf("Error loading brokerpaks: %v", err)
	}

	config, err := config.Parse()
	if err != nil {
		return nil, fmt.Errorf("Failed loading config: %v", err)
	}

	var cs credstore.CredStore

	if config.CredStoreConfig.HasCredHubConfig() {
		var err error
		cs, err = credstore.NewCredhubStore( &config.CredStoreConfig, logger )
		if err != nil {
			return nil, fmt.Errorf("Failed creating credstore: %v", err)
		}
	}

	return &BrokerConfig{
		Registry:   registry,
		Credstore:  cs,
	}, nil
}
