package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"microservice/Models/user"

	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestHandler_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockservice(ctrl)
	h := New(mockService)

	testCases := []struct {
		desc      string
		input     user.User
		expectMsg string
		expectErr bool
		mockErr   error
		status    int
	}{
		{
			desc:      "Valid User",
			input:     user.User{ID: 1, Name: "John", Phone: "1234567890", Email: "john@example.com"},
			expectMsg: "user inserted",
			expectErr: false,
			mockErr:   nil,
			status:    http.StatusOK,
		},
		{
			desc:      "Empty Name",
			input:     user.User{ID: 2, Name: "", Phone: "9876543210", Email: "empty@example.com"},
			expectMsg: "error to insert",
			expectErr: true,
			mockErr:   errors.New("error to insert"),
			status:    http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			jsonData, _ := json.Marshal(tc.input)

			mockService.EXPECT().InsertUser(tc.input).Return(tc.expectMsg, tc.mockErr)

			r := httptest.NewRequest("POST", "/user", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()

			h.AddUser(w, r)

			if w.Code != tc.status {
				t.Errorf("Expected status %d, got %d", tc.status, w.Code)
			}
			body := w.Body.String()

			if !tc.expectErr {

				expectedJSON, _ := json.Marshal(map[string]string{"message": tc.expectMsg})
				if body != string(expectedJSON)+"\n" && body != string(expectedJSON) {
					t.Errorf("Expected body %s, got %s", expectedJSON, body)
				}
			} else {
				if !bytes.Contains([]byte(body), []byte(tc.expectMsg)) {
					t.Errorf("Expected error message to contain %q, got %q", tc.expectMsg, body)
				}
			}
		})
	}
}

func TestHandler_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	t.Run("Valid ID", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(1).Return(&user.User{ID: 1, Name: "Alice"}, nil)
		r := httptest.NewRequest("GET", "/user/1", nil)
		r.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.GetUserByID(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("Expected 200, got %d", w.Code)
		}
	})

	t.Run("Invalid ID Format", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/user/abc", nil)
		r.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.GetUserByID(w, r)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected 400, got %d", w.Code)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(999).Return(nil, errors.New("not found"))
		r := httptest.NewRequest("GET", "/user/999", nil)
		r.SetPathValue("id", "999")
		w := httptest.NewRecorder()
		h.GetUserByID(w, r)
		if w.Code != http.StatusNotFound {
			t.Errorf("Expected 404, got %d", w.Code)
		}
	})
}

func TestHandler_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	mockUsers := []user.User{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}
	mockService.EXPECT().GetAllUsers().Return(mockUsers, nil)

	r := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	h.GetAllUsers(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestHandler_DeleteUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	t.Run("Valid ID", func(t *testing.T) {
		mockService.EXPECT().DeleteUserByID(1).Return("deleted", nil)
		r := httptest.NewRequest("DELETE", "/user/1", nil)
		r.SetPathValue("id", "1")
		w := httptest.NewRecorder()
		h.DeleteUserByID(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("Expected 200, got %d", w.Code)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/user/abc", nil)
		r.SetPathValue("id", "abc")
		w := httptest.NewRecorder()
		h.DeleteUserByID(w, r)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected 400, got %d", w.Code)
		}
	})

	t.Run("Delete Error", func(t *testing.T) {
		mockService.EXPECT().DeleteUserByID(2).Return("", errors.New("delete error"))
		r := httptest.NewRequest("DELETE", "/user/2", nil)
		r.SetPathValue("id", "2")
		w := httptest.NewRecorder()
		h.DeleteUserByID(w, r)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected 500, got %d", w.Code)
		}
	})
}

func TestHandler_DeleteAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockservice(ctrl)
	h := New(mockService)

	mockService.EXPECT().DeleteAllUsers().Return("all deleted", nil)

	r := httptest.NewRequest("DELETE", "/users", nil)
	w := httptest.NewRecorder()
	h.DeleteAllUsers(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
