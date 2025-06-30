package task

import (
	"errors"
	Task_Model "microservice/Models/task"
)

// create a interface of store here that we rhave implement in this code
// Store interface defines all the methods the Service layer depends on.
// This allows the Service to interact with any implementation (real DB, mock, etc.),
// as long as it satisfies this contract.
// It helps achieve decoupling and testability.

// Service struct is the business logic layer.
// It holds a reference to a Store interface, so it can call DB-related methods indirectly.
// This ensures that the Service doesn’t depend on any specific DB implementation.
type Service struct {
	store Store
}

// New is a constructor function for the Service.
// It takes a Store implementation (could be real DB or a mock for testing)
// and returns a pointer to a new Service instance.
func New(s Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) Insertask(t Task_Model.Task) (string, error) {
	if t.Name == "" || t.Status == "" {
		return "", errors.New("task name and status cannot be empty")
	}
	return s.store.Insertask(t)

}
func (s *Service) Getalltask() ([]Task_Model.Task, error) {
	tasks, err := s.store.Getalltask()
	if err != nil {
		return nil, err
	}
	var filtered []Task_Model.Task
	for _, t := range tasks {
		if t.Name != "" {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil

}
func (s *Service) Gettaskbyid(id int) (*Task_Model.Task, error) {
	if id <= 0 {
		return nil, errors.New("id should not be negative")
	}
	return s.store.Gettaskbyid(id)

}
func (s *Service) Deletetask(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("id should not be negative")
	}
	return s.store.Deletetask(id)

}
func (s *Service) Completetask(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("id should not be negative")
	}
	return s.store.Completetask(id)

}
