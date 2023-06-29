package customer

import (
	"fmt"
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

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) Create(model Model) (string, error) {
	if model.Address == "" || model.Contact == "" || model.Name == "" {
		return "", fmt.Errorf("required field(s) missing")
	}
	return s.repository.Create(model)
}
