package task

import (
	"errors"
	models_task "microservice/Models/task"
)

// create a interface of store here that we rhave implement in this code
// Store interface defines all the methods the Service layer depends on.
// This allows the Service to interact with any implementation (real DB, mock, etc.),
// as long as it satisfies this contract.
// It helps achieve decoupling and testability.
type Store interface {
	Insertask(t models_task.Task) (string, error)
	Getalltask() ([]models_task.Task, error)
	Gettaskbyid(id int) (*models_task.Task, error)
	Deletetask(id int) (string, error)
	Completetask(id int) (string, error)
}

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

func (s *Service) Insertask(t models_task.Task) (string, error) {
	if t.TaskName == "" || t.TaskStatus == "" {
		return "", errors.New("task name and status cannot be empty")
	}
	return s.store.Insertask(t)

}
func (s *Service) Getalltask() ([]models_task.Task, error) {
	tasks, err := s.store.Getalltask()
	if err != nil {
		return nil, err
	}
	var filtered []models_task.Task
	for _, t := range tasks {
		if t.TaskName != "" {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil

}
func (s *Service) Gettaskbyid(id int) (*models_task.Task, error) {
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
