package task

import models_task "microservice/Models/task"

type service interface {
	Insertask(t models_task.Task) (string, error)
	Getalltask() ([]models_task.Task, error)
	Gettaskbyid(id int) (*models_task.Task, error)
	Deletetask(id int) (string, error)
	Completetask(id int) (string, error)
}
