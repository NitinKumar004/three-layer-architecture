package task

import (
	"errors"
	"fmt"
	"microservice/Models/task"
	"testing"
)

func InsertTaskFunc(t task.Task) (string, error) {
	return "task inserted", nil
}

func GetalltaskFunc() ([]task.Task, error) {
	alltask := []task.Task{
		task.Task{
			TaskID:     1,
			TaskName:   "complete test cover",
			TaskStatus: "pending",
			AssignUser: 1,
		}, {
			TaskID:     21,
			TaskName:   "complete test cover to 100%",
			TaskStatus: "complete",
			AssignUser: 64,
		},
	}
	return alltask, nil

}

func GettaskbyidFunc(id int) (*task.Task, error) {

	D := task.Task{
		TaskID:     1,
		TaskName:   "complete test cover",
		TaskStatus: "pending",
		AssignUser: 1,
	}

	return &D, nil

}
func TestInsertask_Success(t *testing.T) {
	fmt.Println("----------------")
	fmt.Println(" Testing of USER WITH SERVICE  ")
	fmt.Println("----------------")
	mock := MockStore{
		InsertTaskFunc: InsertTaskFunc,
	}
	s := New(&mock)
	task := task.Task{
		TaskID:     1,
		TaskName:   "testing file",
		TaskStatus: "pending",
		AssignUser: 4,
	}
	data, err := s.Insertask(task)
	if err != nil || data != "task inserted" {
		t.Errorf("Expected 'task inserted', got '%s', error: %v", data, err)
	}

}
func TestGETBYID_Success(t *testing.T) {
	mock := MockStore{
		GettaskbyidFunc: GettaskbyidFunc,
	}
	s := New(&mock)

	data, err := s.Gettaskbyid(1)
	if err != nil {
		t.Errorf("errror to getting call")
	}
	if data.TaskName != "complete test cover" || data.TaskStatus != "pending" || data.TaskID != 1 || data.AssignUser != 1 {
		t.Errorf("getting task data. Got: TaskName=%s, TaskStatus=%s, TaskID=%d, AssignUser=%d",
			data.TaskName, data.TaskStatus, data.TaskID, data.AssignUser)
	}

}
func TestMockStore_Getalltask(t *testing.T) {
	mock := MockStore{
		GetalltaskFunc: GetalltaskFunc,
	}
	s := New(&mock)
	data, err := s.Getalltask()
	if err != nil {
		t.Errorf("errror to getting call")
	}
	n := len(data)
	if n != 2 {
		t.Errorf("expected %d and got this %d", 2, n)
	}

}
func TestService_Deletetask(t *testing.T) {
	mock := MockStore{
		DeletetaskFunc: func(id int) (string, error) {
			return "delete task successfully", nil
		},
	}

	s := New(&mock)
	d, err := s.Deletetask(1)
	if err != nil {
		t.Errorf("errror to getting call")
	}
	if d != "delete task successfully" {
		t.Errorf("expected this %s and got this %s ", "delete task successfully", d)
	}
}

func TestService_Completetask(t *testing.T) {
	mock := MockStore{
		CompletetaskFunc: func(id int) (string, error) {
			return "complete task successfully", nil
		},
	}
	s := New(&mock)
	d, err := s.Completetask(1)
	if err != nil {
		t.Errorf("errror to getting call")
	}
	if d != "complete task successfully" {
		t.Errorf("expected this %s and got this %s", "complete task successfully", d)
	}

}
func TestInsertask_Error(t *testing.T) {
	mock := MockStore{
		InsertTaskFunc: func(t task.Task) (string, error) {
			return "", errors.New("insert error")
		},
	}
	s := New(&mock)

	_, err := s.Insertask(task.Task{
		TaskID:     99,
		TaskName:   "fail",
		TaskStatus: "error",
		AssignUser: 0,
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGETBYID_Error(t *testing.T) {
	mock := MockStore{
		GettaskbyidFunc: func(id int) (*task.Task, error) {
			return nil, errors.New("not found")
		},
	}
	s := New(&mock)

	data, err := s.Gettaskbyid(999)
	if err == nil || data != nil {
		t.Errorf("Expected error and nil task, got %+v, err: %v", data, err)
	}
}

func TestGetalltask_Error(t *testing.T) {
	mock := MockStore{
		GetalltaskFunc: func() ([]task.Task, error) {
			return nil, errors.New("database error")
		},
	}
	s := New(&mock)

	data, err := s.Getalltask()
	if err == nil || data != nil {
		t.Errorf("Expected error and nil slice, got %+v, err: %v", data, err)
	}
}

func TestDeletetask_Error(t *testing.T) {
	mock := MockStore{
		DeletetaskFunc: func(id int) (string, error) {
			return "", errors.New("delete failed")
		},
	}
	s := New(&mock)

	msg, err := s.Deletetask(404)
	if err == nil || msg != "" {
		t.Errorf("Expected error and empty message, got '%s', err: %v", msg, err)
	}
}

func TestCompletetask_Error(t *testing.T) {
	mock := MockStore{
		CompletetaskFunc: func(id int) (string, error) {
			return "", errors.New("complete failed")
		},
	}
	s := New(&mock)

	msg, err := s.Completetask(777)
	if err == nil || msg != "" {
		t.Errorf("Expected error and empty message, got '%s', err: %v", msg, err)
	}
}
