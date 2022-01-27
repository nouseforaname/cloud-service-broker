package storage

import (
	"fmt"

	"github.com/cloudfoundry-incubator/cloud-service-broker/db_service/models"
)

type TerraformDeployment struct {
	ID                   string
	Workspace            []byte
	LastOperationType    string
	LastOperationState   string
	LastOperationMessage string
}

func (s *Storage) StoreTerraformDeployment(t TerraformDeployment) error {
	encoded, err := s.encodeBytes(t.Workspace)
	if err != nil {
		return fmt.Errorf("error encoding workspace: %w", err)
	}

	var m models.TerraformDeployment
	if err := s.loadTerraformDeploymentIfExists(t.ID, &m); err != nil {
		return err
	}

	m.Workspace = encoded
	m.LastOperationType = t.LastOperationType
	m.LastOperationState = t.LastOperationState
	m.LastOperationMessage = t.LastOperationMessage

	switch m.ID {
	case "":
		m.ID = t.ID
		if err := s.db.Create(&m).Error; err != nil {
			return fmt.Errorf("error creating terraform deployment: %w", err)
		}
	default:
		if err := s.db.Save(&m).Error; err != nil {
			return fmt.Errorf("error saving terraform deployment: %w", err)
		}
	}

	return nil
}

func (s *Storage) GetTerraformDeployment(id string) (TerraformDeployment, error) {
	exists, err := s.ExistsTerraformDeployment(id)
	switch {
	case err != nil:
		return TerraformDeployment{}, err
	case !exists:
		return TerraformDeployment{}, fmt.Errorf("could not find terraform deployment: %s", id)
	}

	var receiver models.TerraformDeployment
	if err := s.db.Where("id = ?", id).First(&receiver).Error; err != nil {
		return TerraformDeployment{}, fmt.Errorf("error finding terraform deployment: %w", err)
	}

	decoded, err := s.decodeBytes(receiver.Workspace)
	if err != nil {
		return TerraformDeployment{}, fmt.Errorf("error decoding workspace: %w", err)
	}

	return TerraformDeployment{
		ID:                   id,
		Workspace:            decoded,
		LastOperationType:    receiver.LastOperationType,
		LastOperationState:   receiver.LastOperationState,
		LastOperationMessage: receiver.LastOperationMessage,
	}, nil
}

func (s *Storage) ExistsTerraformDeployment(id string) (bool, error) {
	var count int64
	if err := s.db.Model(&models.TerraformDeployment{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, fmt.Errorf("error counting terraform deployments: %w", err)
	}
	return count != 0, nil
}

func (s *Storage) DeleteTerraformDeployment(id string) error {
	err := s.db.Where("id = ?", id).Delete(&models.TerraformDeployment{}).Error
	if err != nil {
		return fmt.Errorf("error deleting terraform deployment: %w", err)
	}
	return nil
}

func (s *Storage) loadTerraformDeploymentIfExists(id string, receiver interface{}) error {
	if id == "" {
		return nil
	}

	exists, err := s.ExistsTerraformDeployment(id)
	switch {
	case err != nil:
		return err
	case !exists:
		return nil
	}

	return s.db.Where("id = ?", id).First(receiver).Error
}