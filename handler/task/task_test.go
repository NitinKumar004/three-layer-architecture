package task

import (
	"bytes"
	"encoding/json"
	"errors"

	"microservice/Models/task"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestHandler_Addtask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockservice(ctrl)
	h := New(mockService)

	testCases := []struct {
		desc      string
		input     task.Task
		expectMsg string
		expectErr bool
	}{
		{"Valid Task", task.Task{ID: 1, Name: "Task One", Status: "pending", UserID: 10}, "task inserted", false},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.input)

			if !tc.expectErr {
				mockService.EXPECT().Insertask(tc.input).Return(tc.expectMsg, nil)
			}

			r := httptest.NewRequest("POST", "/task", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()

			h.Addtask(w, r)

			if tc.expectErr && w.Code != http.StatusInternalServerError && w.Code != http.StatusBadRequest {
				t.Errorf("Expected error status, got %d", w.Code)
			}
			if !tc.expectErr && w.Code != http.StatusCreated {
				t.Errorf("Expected 201, got %d", w.Code)
			}
		})
	}
}
func TestHandler_Gettaskbyid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	testCases := []struct {
		desc      string
		id        string
		mockID    int
		mockResp  *task.Task
		expectErr bool
		code      int
	}{
		{"Valid ID", "1", 1, &task.Task{ID: 1, Name: "Mock Task"}, false, http.StatusOK},
		{"Invalid ID Format", "abc", 0, nil, true, http.StatusInternalServerError},
		{"Not Found", "999", 999, nil, true, http.StatusNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/task/"+tc.id, nil)
			r.SetPathValue("id", tc.id)
			w := httptest.NewRecorder()

			if !tc.expectErr {
				mockService.EXPECT().Gettaskbyid(tc.mockID).Return(tc.mockResp, nil)
			} else if tc.code == http.StatusNotFound {
				mockService.EXPECT().Gettaskbyid(tc.mockID).Return(nil, errors.New("not found"))
			}

			h.Gettaskbyid(w, r)
			if w.Code != tc.code {
				t.Errorf("Expected status %d, got %d", tc.code, w.Code)
			}
		})
	}
}

func TestHandler_Getalltask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	tasks := []task.Task{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
	}

	mockService.EXPECT().Getalltask().Return(tasks, nil)

	r := httptest.NewRequest("GET", "/task", nil)
	w := httptest.NewRecorder()

	h.Getalltask(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestHandler_Deletetask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	testCases := []struct {
		desc      string
		id        string
		mockID    int
		mockResp  string
		expectErr bool
		code      int
	}{
		{"Valid ID", "1", 1, "deleted", false, http.StatusOK},
		{"Invalid ID Format", "abc", 0, "", true, http.StatusBadRequest},
		{"Delete Error", "2", 2, "", true, http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := httptest.NewRequest("DELETE", "/task/"+tc.id, nil)
			r.SetPathValue("id", tc.id)
			w := httptest.NewRecorder()

			if !tc.expectErr {
				mockService.EXPECT().Deletetask(tc.mockID).Return(tc.mockResp, nil)
			} else if tc.code == http.StatusInternalServerError {
				mockService.EXPECT().Deletetask(tc.mockID).Return("", errors.New("delete error"))
			}

			h.Deletetask(w, r)
			if w.Code != tc.code {
				t.Errorf("Expected %d, got %d", tc.code, w.Code)
			}
		})
	}
}

func TestHandler_Completetask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	testCases := []struct {
		desc      string
		id        string
		mockID    int
		mockResp  string
		expectErr bool
		code      int
	}{
		{"Valid ID", "1", 1, "completed", false, http.StatusOK},
		{"Invalid ID Format", "xyz", 0, "", true, http.StatusBadRequest},
		{"Error Completing", "2", 2, "", true, http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := httptest.NewRequest("PATCH", "/task/"+tc.id, nil)
			r.SetPathValue("id", tc.id)
			w := httptest.NewRecorder()

			if !tc.expectErr {
				mockService.EXPECT().Completetask(tc.mockID).Return(tc.mockResp, nil)
			} else if tc.code == http.StatusInternalServerError {
				mockService.EXPECT().Completetask(tc.mockID).Return("", errors.New("completion error"))
			}

			h.Completetask(w, r)
			if w.Code != tc.code {
				t.Errorf("Expected %d, got %d", tc.code, w.Code)
			}
		})
	}
}
