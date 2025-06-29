package user

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	usermodel "microservice/Models/user"
	"testing"
)

func TestStore_InsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error to connecting fake db connections %v", err)
	}
	s := New(db)
	newuser := usermodel.User{
		UserID:    1,
		UserName:  "nitin patel",
		UserPhone: "8558856856",
		UserEmail: "nitinraj7488@gmail.com",
	}
	mock.ExpectExec("INSERT INTO usermanage").WithArgs(newuser.UserID, newuser.UserName, newuser.UserPhone, newuser.UserEmail).WillReturnResult(sqlmock.NewResult(1, 1))
	msg, err := s.InsertUser(newuser)
	if msg != "user inserted successfully" {
		t.Errorf("expected this %s and got this %s", "user inserted successfully", msg)

	}
	if mock.ExpectationsWereMet() != nil {
		t.Errorf("there are something that we have missed %v", err)
	}

}
func TestStore_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error to connecting fake db connections %v", err)
	}
	s := New(db)
	newuser := usermodel.User{
		UserID:    1,
		UserName:  "nitin patel",
		UserPhone: "8558856856",
		UserEmail: "nitinraj7488@gmail.com",
	}
	TASKID := 1
	mock.ExpectQuery("SELECT userid, username, userphone, useremail FROM usermanage WHERE userid = ?").WithArgs(TASKID).WillReturnRows(sqlmock.NewRows([]string{"userid", "username", "userphone", "useremail"}).AddRow(newuser.UserID, newuser.UserName, newuser.UserPhone, newuser.UserEmail))
	actual, _ := s.GetUserByID(TASKID)
	if actual.UserID != newuser.UserID || actual.UserEmail != newuser.UserEmail || actual.UserPhone != newuser.UserPhone || actual.UserName != newuser.UserName {
		t.Errorf("Expected %+v, got %+v", newuser, actual)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
func TestStore_GetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	s := New(db)

	expectedUsers := []usermodel.User{
		{UserID: 1, UserName: "Nitin", UserPhone: "1234567890", UserEmail: "nitin@example.com"},
		{UserID: 2, UserName: "Patel", UserPhone: "9876543210", UserEmail: "patel@example.com"},
	}

	rows := sqlmock.NewRows([]string{"userid", "username", "userphone", "useremail"}).
		AddRow(expectedUsers[0].UserID, expectedUsers[0].UserName, expectedUsers[0].UserPhone, expectedUsers[0].UserEmail).
		AddRow(expectedUsers[1].UserID, expectedUsers[1].UserName, expectedUsers[1].UserPhone, expectedUsers[1].UserEmail)

	mock.ExpectQuery("SELECT userid, username, userphone, useremail FROM usermanage").WillReturnRows(rows)

	users, err := s.GetAllUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(users) != len(expectedUsers) {
		t.Errorf("expected %d users, got %d", len(expectedUsers), len(users))
	}

	for i := range users {
		if users[i] != expectedUsers[i] {
			t.Errorf("expected %+v, got %+v", expectedUsers[i], users[i])
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestStore_DeleteAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	s := New(db)

	mock.ExpectExec("DELETE FROM usermanage").WillReturnResult(sqlmock.NewResult(0, 2))

	msg, err := s.DeleteAllUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "all users deleted successfully" {
		t.Errorf("expected success message, got: %s", msg)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestStore_Deletetaskbyid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	s := New(db)
	userid := 1

	mock.ExpectExec("DELETE FROM usermanage WHERE userid =?").
		WithArgs(userid).
		WillReturnResult(sqlmock.NewResult(0, 1))

	msg, err := s.DeleteUserByID(userid)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if msg != "user deleted successfully" {
		t.Errorf("expected success message, got: %s", msg)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestStore_InsertUser_SQL_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	newuser := usermodel.User{
		UserID:    1,
		UserName:  "error user",
		UserPhone: "0000000000",
		UserEmail: "error@example.com",
	}

	mock.ExpectExec("INSERT INTO usermanage").
		WithArgs(newuser.UserID, newuser.UserName, newuser.UserPhone, newuser.UserEmail).
		WillReturnError(errors.New("primary key violation"))

	msg, err := s.InsertUser(newuser)
	if err == nil {
		t.Error("expected error, got nil")
	}
	fmt.Println(msg)
	if msg != "user not inserted, maybe primary key issue" {
		t.Errorf("unexpected message: %s", msg)
	}
}
func TestStore_GetUserByID_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	mock.ExpectQuery("SELECT userid, username, userphone, useremail FROM usermanage WHERE userid = ?").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows([]string{"userid", "username", "userphone", "useremail"}))

	_, err := s.GetUserByID(999)
	if err == nil {
		t.Error("expected error for missing user, got nil")
	}
}
func TestStore_DeleteAllUsers_NoRowsAffected(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	mock.ExpectExec("DELETE FROM usermanage").
		WillReturnResult(sqlmock.NewResult(0, 0))

	msg, err := s.DeleteAllUsers()
	if msg != "no users found" || err == nil {
		t.Errorf("expected error for no users deleted, got: %s, err: %v", msg, err)
	}
}
