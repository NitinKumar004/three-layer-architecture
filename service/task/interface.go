package task

import models_task "microservice/Models/task"

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
