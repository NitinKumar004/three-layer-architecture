package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	taskmodel "microservice/Models/task"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Addtask(t *testing.T) {
	fmt.Println("----------------")
	fmt.Println(" Testing of TASK WITH HANDLER  ")
	fmt.Println("----------------")
	mock := &Mockstore{
		InsertTaskFunc: func(t taskmodel.Task) (string, error) {
			return "task inserted", nil
		},
	}

	h := New(mock)

	task := taskmodel.Task{
		TaskID:     1,
		TaskName:   "working on project",
		TaskStatus: "pending",
		AssignUser: 2,
	}

	jsontask, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("failed to marshal task: %v", err)
	}

	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/task", bytes.NewBuffer(jsontask))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.Addtask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expected := `{"message":"task inserted"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}

func TestHandler_Gettaskbyid(t *testing.T) {
	mock := &Mockstore{GettaskbyidFunc: func(id int) (*taskmodel.Task, error) {
		task := taskmodel.Task{
			TaskID:     1,
			TaskName:   "GO HOME",
			TaskStatus: "pending",
			AssignUser: 1,
		}
		return &task, nil

	},
	}
	h := New(mock)
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/task/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	req.SetPathValue("id", "1")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Gettaskbyid)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}
	var tasknew taskmodel.Task
	err = json.Unmarshal(rr.Body.Bytes(), &tasknew)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)

	}
	if tasknew.TaskID != 1 && tasknew.TaskName != "GO HOME" {
		t.Errorf("expected this %d %s got this %d %s", 1, "GO HOME", tasknew.TaskID, tasknew.TaskName)
	}

}
func TestHandler_Getalltask(t *testing.T) {
	mock := &Mockstore{GetalltaskFunc: func() ([]taskmodel.Task, error) {
		alltask := []taskmodel.Task{
			{TaskID: 1,
				TaskName:   "GO HOME",
				TaskStatus: "pending",
				AssignUser: 1},
			{TaskID: 2,
				TaskName:   "GO HOME by bus",
				TaskStatus: "complete",
				AssignUser: 44},
		}
		return alltask, nil

	},
	}
	h := New(mock)
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/task", http.NoBody)
	if err != nil {
		t.Fatal(err)

	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Getalltask)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}
	var alltask []taskmodel.Task
	err = json.Unmarshal(rr.Body.Bytes(), &alltask)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)

	}
	if len(alltask) != 2 {
		t.Errorf("exprected this %d and got this %d", 2, len(alltask))
	}
}
func TestHandler_Completetask(t *testing.T) {
	mock := &Mockstore{
		CompletetaskFunc: func(i int) (string, error) {
			return "completed task successfully", nil

		},
	}
	h := New(mock)
	req, err := http.NewRequestWithContext(t.Context(), http.MethodPatch, "/task/", http.NoBody)
	if err != nil {
		return

	}
	req.SetPathValue("id", "1")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Completetask)
	handler.ServeHTTP(rr, req)

	expected := `{"message":"completed task successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("expected this got this  %s and got this %s", expected, rr.Body.String())
	}

}
func TestHandler_Deletetask(t *testing.T) {
	mock := &Mockstore{
		DeletetaskFunc: func(i int) (string, error) {
			return "delete task successfully", nil

		},
	}
	h := New(mock)
	req, err := http.NewRequestWithContext(t.Context(), http.MethodDelete, "/task/", http.NoBody)
	if err != nil {
		return
	}
	req.SetPathValue("id", "1")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Deletetask)
	handler.ServeHTTP(rr, req)

	expected := `{"message":"delete task successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("expected this got this  %s and got this %s", expected, rr.Body.String())
	}

}
func TestHandler_Addtask_Error(t *testing.T) {
	mock := &Mockstore{
		InsertTaskFunc: func(t taskmodel.Task) (string, error) {
			return "", fmt.Errorf("failed to insert")
		},
	}

	h := New(mock)

	task := taskmodel.Task{
		TaskID:     1,
		TaskName:   "broken",
		TaskStatus: "pending",
		AssignUser: 2,
	}

	jsontask, _ := json.Marshal(task)
	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(jsontask))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.Addtask).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestHandler_Gettaskbyid_Error(t *testing.T) {
	mock := &Mockstore{
		GettaskbyidFunc: func(id int) (*taskmodel.Task, error) {
			return nil, errors.New("not found")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodGet, "/task/", http.NoBody)
	req.SetPathValue("id", "1")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.Gettaskbyid).ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestHandler_Getalltask_Error(t *testing.T) {
	mock := &Mockstore{
		GetalltaskFunc: func() ([]taskmodel.Task, error) {
			return nil, errors.New("db failure")
		},
	}

	h := New(mock)
	req := httptest.NewRequest(http.MethodGet, "/task", http.NoBody)
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.Getalltask).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestHandler_Completetask_Error(t *testing.T) {
	mock := &Mockstore{
		CompletetaskFunc: func(id int) (string, error) {
			return "", fmt.Errorf("fail to complete")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodPatch, "/task/", http.NoBody)
	req.SetPathValue("id", "1")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.Completetask).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}

func TestHandler_Deletetask_Error(t *testing.T) {
	mock := &Mockstore{
		DeletetaskFunc: func(id int) (string, error) {
			return "", fmt.Errorf("failed to delete")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodDelete, "/task/", http.NoBody)
	req.SetPathValue("id", "123")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.Deletetask).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}
}
