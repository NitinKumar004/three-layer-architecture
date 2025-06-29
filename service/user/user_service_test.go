package user

import (
	"errors"
	"fmt"
	models_user "microservice/Models/user"
	"testing"
)

func TestMockstore_InsertUser(t *testing.T) {
	fmt.Println("----------------")
	fmt.Println(" Testing of USER WITH Service  ")
	fmt.Println("----------------")
	mock := Mockstore{Insertnewuserfunc: func(user models_user.User) (string, error) {
		return "Insert new user successfully", nil

	},
	}
	s := New(&mock)
	newuser := models_user.User{
		UserID:    12,
		UserEmail: "nitinraj7488204975@gmail.com",
		UserName:  "Nitin Kumar",
		UserPhone: "7488204975",
	}
	data, err := s.InsertUser(newuser)
	if err != nil || data != "Insert new user successfully" {
		t.Errorf("expected this %s and got this %s", "Insert new user successfully", data)
	}

}

func TestMockstore_GetUserByID(t *testing.T) {
	mock := Mockstore{GetUserByIDfunc: func(id int) (*models_user.User, error) {
		task := models_user.User{
			UserID:    12,
			UserEmail: "nitinraj7488204975@gmail.com",
			UserName:  "Nitin Kumar",
			UserPhone: "7488204975",
		}
		return &task, nil

	},
	}
	s := New(&mock)
	data, err := s.GetUserByID(1)
	if err != nil {
		t.Errorf("error to fetching the data")
	}
	if data.UserID != 12 || data.UserEmail != "nitinraj7488204975@gmail.com" || data.UserName != "Nitin Kumar" || data.UserPhone != "7488204975" {
		t.Errorf("expected this %d %s %s %s and got this %d %s %s %s", 12, "nitinraj7488204975@gmail.com", "Nitin Kumar", "7488204975",
			data.UserID, data.UserEmail, data.UserName, data.UserPhone)
	}

}

func TestMockstore_GetAllUsers(t *testing.T) {
	mock := Mockstore{Getalluserfunc: func() ([]models_user.User, error) {
		data := []models_user.User{
			models_user.User{
				UserID:    1,
				UserName:  "nitin",
				UserEmail: "nitin@gmail.com",
				UserPhone: "7488204975",
			},
			models_user.User{
				UserID:    2,
				UserName:  "nishant",
				UserEmail: "nishant@gmail.com",
				UserPhone: "248355335",
			},
		}
		return data, nil

	}}
	s := New(&mock)
	d, err := s.GetAllUsers()
	if err != nil || len(d) != 2 {
		t.Errorf("expected this %d and got this %d", 2, len(d))
	}

}
func TestService_DeleteUserByID(t *testing.T) {
	mock := Mockstore{
		DeleteUserbyidfunc: func(id int) (string, error) {
			return "delete successfully", nil
		},
	}
	s := New(&mock)
	msg, err := s.DeleteUserByID(1)
	if err != nil || msg != "delete successfully" {
		t.Errorf("expected this %s and got this %s", "delete successfully", msg)

	}

}
func TestService_DeleteAllUsers(t *testing.T) {
	mock := Mockstore{
		Deletealluserfunc: func() (string, error) {
			return "delete all  successfully", nil
		},
	}
	s := New(&mock)
	msg, err := s.DeleteAllUsers()
	if err != nil || msg != "delete all  successfully" {
		t.Errorf("expected this %s and got this %s", "delete all  successfully", msg)

	}

}
func TestService_InsertUser_Error(t *testing.T) {
	mock := Mockstore{
		Insertnewuserfunc: func(user models_user.User) (string, error) {
			return "", errors.New("insert failed")
		},
	}
	s := New(&mock)

	newuser := models_user.User{
		UserID:    99,
		UserEmail: "fail@example.com",
		UserName:  "Fail Case",
		UserPhone: "0000000000",
	}
	msg, err := s.InsertUser(newuser)
	if err == nil || msg != "" {
		t.Errorf("Expected error, got message: '%s', error: %v", msg, err)
	}
}

func TestService_GetUserByID_Error(t *testing.T) {
	mock := Mockstore{
		GetUserByIDfunc: func(id int) (*models_user.User, error) {
			return nil, errors.New("user not found")
		},
	}
	s := New(&mock)

	user, err := s.GetUserByID(404)
	if err == nil || user != nil {
		t.Errorf("Expected error and nil user, got: %+v, err: %v", user, err)
	}
}

func TestService_GetAllUsers_Error(t *testing.T) {
	mock := Mockstore{
		Getalluserfunc: func() ([]models_user.User, error) {
			return nil, errors.New("failed to fetch users")
		},
	}
	s := New(&mock)

	users, err := s.GetAllUsers()
	if err == nil || users != nil {
		t.Errorf("Expected error and nil slice, got: %+v, err: %v", users, err)
	}
}

func TestService_DeleteUserByID_Error(t *testing.T) {
	mock := Mockstore{
		DeleteUserbyidfunc: func(id int) (string, error) {
			return "", errors.New("delete failed")
		},
	}
	s := New(&mock)

	msg, err := s.DeleteUserByID(777)
	if err == nil || msg != "" {
		t.Errorf("Expected error and empty message, got: '%s', err: %v", msg, err)
	}
}

func TestService_DeleteAllUsers_Error(t *testing.T) {
	mock := Mockstore{
		Deletealluserfunc: func() (string, error) {
			return "", errors.New("delete all failed")
		},
	}
	s := New(&mock)

	msg, err := s.DeleteAllUsers()
	if err == nil || msg != "" {
		t.Errorf("Expected error and empty message, got: '%s', err: %v", msg, err)
	}
}
