package customer

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func (s *Service) GetById(uuid string) (Model, error) {
	return s.repository.GetById(uuid)
}

func (s *Service) GetAll() ([]Model, error) {
	return s.repository.GetAll()
}

func (s *Service) Delete(uuid string) error {
	return s.repository.Delete(uuid)
}

func (s *Service) Create(model Model) (string, error) {
	//@TODO: Validate if uuid already exists in our database
	model.Uuid = uuid.New().String()
	if model.Address == "" || model.Contact == "" || model.Name == "" {
		return "", fmt.Errorf("required field(s) missing")
	}
	return s.repository.Create(model)
}
