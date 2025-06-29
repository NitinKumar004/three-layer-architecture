package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	models_user "microservice/Models/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_AddUser(t *testing.T) {
	fmt.Println("----------------")
	fmt.Println(" Testing of USER WITH HANDLER  ")
	fmt.Println("----------------")
	mock := &MockStore{
		Insertnewuserfunc: func(u models_user.User) (string, error) {
			return "user inserted", nil
		},
	}
	h := New(mock)

	user := models_user.User{
		UserID:    1,
		UserName:  "NITIN",
		UserPhone: "7488204975",
		UserEmail: "nit@NITIN.com",
	}
	jsonuser, _ := json.Marshal(user)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/user", bytes.NewBuffer(jsonuser))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.AddUser)
	handler.ServeHTTP(rr, req)

	expected := `{"message":"user inserted"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected %s, got %s", expected, rr.Body.String())
	}
}

func TestHandler_GetUserByID(t *testing.T) {
	mock := &MockStore{
		GetUserByIDfunc: func(id int) (*models_user.User, error) {
			return &models_user.User{
				UserID:    1,
				UserName:  "nitin",
				UserPhone: "7488204975",
				UserEmail: "nitin@example.com",
			}, nil
		},
	}
	h := New(mock)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/user", nil)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetUserByID)
	handler.ServeHTTP(rr, req)

	var user models_user.User
	err := json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil || user.UserID != 1 {
		t.Errorf("Expected user ID 1, got %+v, err: %v", user, err)
	}
}

// --- Test: GetAllUsers ---
func TestHandler_GetAllUsers(t *testing.T) {
	mock := &MockStore{
		Getalluserfunc: func() ([]models_user.User, error) {
			return []models_user.User{
				{UserID: 1, UserName: "A"},
				{UserID: 2, UserName: "B"},
			}, nil
		},
	}
	h := New(mock)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.GetAllUsers)
	handler.ServeHTTP(rr, req)

	var users []models_user.User
	err := json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil || len(users) != 2 {
		t.Errorf("Expected 2 users, got %d, err: %v", len(users), err)
	}
}

// --- Test: DeleteUserByID ---
func TestHandler_DeleteUserByID(t *testing.T) {
	mock := &MockStore{
		DeleteUserbyidfunc: func(id int) (string, error) {
			if id == 1 {
				return "user deleted", nil
			}
			return "", nil
		},
	}
	h := New(mock)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodDelete, "/user", nil)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeleteUserByID)
	handler.ServeHTTP(rr, req)

	expected := `{"message":"user deleted"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected %s, got %s", expected, rr.Body.String())
	}
}

// --- Test: DeleteAllUsers ---
func TestHandler_DeleteAllUsers(t *testing.T) {
	mock := &MockStore{
		Deletealluserfunc: func() (string, error) {
			return "all users deleted", nil
		},
	}
	h := New(mock)

	req := httptest.NewRequestWithContext(t.Context(), http.MethodDelete, "/users", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.DeleteAllUsers)
	handler.ServeHTTP(rr, req)

	expected := `{"message":"all users deleted"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected %s, got %s", expected, rr.Body.String())
	}
}
func TestHandler_AddUser_Error(t *testing.T) {
	mock := &MockStore{
		Insertnewuserfunc: func(u models_user.User) (string, error) {
			return "", errors.New("insert failed")
		},
	}
	h := New(mock)
	user := models_user.User{UserID: 2}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.AddUser).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rr.Code)
	}
}

func TestHandler_GetUserByID_Error(t *testing.T) {
	mock := &MockStore{
		GetUserByIDfunc: func(id int) (*models_user.User, error) {
			return nil, errors.New("user not found")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.SetPathValue("id", "404")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.GetUserByID).ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 500, got %d", rr.Code)
	}
}

func TestHandler_GetAllUsers_Error(t *testing.T) {
	mock := &MockStore{
		Getalluserfunc: func() ([]models_user.User, error) {
			return nil, errors.New("fetch error")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.GetAllUsers).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rr.Code)
	}
}

func TestHandler_DeleteUserByID_Error(t *testing.T) {
	mock := &MockStore{
		DeleteUserbyidfunc: func(id int) (string, error) {
			return "", errors.New("delete failed")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodDelete, "/user", nil)
	req.SetPathValue("id", "99")
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.DeleteUserByID).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rr.Code)
	}
}

func TestHandler_DeleteAllUsers_Error(t *testing.T) {
	mock := &MockStore{
		Deletealluserfunc: func() (string, error) {
			return "", errors.New("delete all failed")
		},
	}
	h := New(mock)

	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rr := httptest.NewRecorder()

	http.HandlerFunc(h.DeleteAllUsers).ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rr.Code)
	}
}
