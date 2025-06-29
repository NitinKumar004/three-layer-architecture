package task

import (
	models_task "microservice/Models/task"
)

type Mockstore struct {
	InsertTaskFunc   func(models_task.Task) (string, error)
	GettaskbyidFunc  func(int) (*models_task.Task, error)
	GetalltaskFunc   func() ([]models_task.Task, error)
	DeletetaskFunc   func(int) (string, error)
	CompletetaskFunc func(int) (string, error)
}

func (m *Mockstore) Insertask(t models_task.Task) (string, error) {
	return m.InsertTaskFunc(t)

}
func (m *Mockstore) Getalltask() ([]models_task.Task, error) {
	return m.GetalltaskFunc()
}
func (m *Mockstore) Gettaskbyid(id int) (*models_task.Task, error) {
	return m.GettaskbyidFunc(id)

}
func (m *Mockstore) Deletetask(id int) (string, error) {
	return m.DeletetaskFunc(id)
}

func (m *Mockstore) Completetask(id int) (string, error) {
	return m.CompletetaskFunc(id)
}
