package task

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	models "microservice/Models/task"
	"testing"
)

func getMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	return sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
}

func TestStore_Insertask(t *testing.T) {

	db, mock, err := getMockDB(t)
	if err != nil {
		t.Errorf("error to connecting fake db connections %v", err)

	}
	s := New(db)
	task := models.Task{
		ID:     1,
		Name:   "Test Task",
		Status: "pending",
		UserID: 99,
	}
	mock.ExpectExec("INSERT INTO taskmanage(taskid, taskname, status, assigned_user_id)VALUES(?,?,?,?)").WithArgs(task.ID, task.Name, task.Status, task.UserID).WillReturnResult(sqlmock.NewResult(1, 1))
	msg, _ := s.Insertask(task)
	if msg != "data inserted successfully" {
		t.Errorf("expected success message, got: %s", msg)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there are something that we have missed %v", err)
	}
}
func TestStore_Insertask_Error(t *testing.T) {
	db, mock, _ := getMockDB(t)

	s := New(db)

	task := models.Task{ID: 1, Name: "Test Task", Status: "pending", UserID: 10}

	mock.ExpectExec("INSERT INTO taskmanage").
		WithArgs(task.ID, task.Name, task.Status, task.UserID).
		WillReturnError(errors.New("primary key violation"))

	msg, err := s.Insertask(task)
	if err == nil || msg != "data not inserted may be primary key issue here is issue : " {
		t.Errorf("expected error msg, got: %s, err: %v", msg, err)
	}
}
func TestStore_Gettaskbyid(t *testing.T) {
	db, mock, err := getMockDB(t)
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	s := New(db)

	expected := models.Task{
		ID:     1,
		Name:   "Test Task",
		Status: "pending",
		UserID: 99,
	}

	mock.ExpectQuery("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage  where taskid=?").
		WithArgs(expected.ID).
		WillReturnRows(sqlmock.NewRows([]string{"taskid", "taskname", "status", "assigned_user_id"}).
			AddRow(expected.ID, expected.Name, expected.Status, expected.UserID))

	actual, err := s.Gettaskbyid(expected.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual.ID != expected.ID || actual.Name != expected.Name ||
		actual.Status != expected.Status || actual.UserID != expected.UserID {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
func TestStore_Gettaskbyid_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	mock.ExpectQuery("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows([]string{"taskid", "taskname", "status", "assigned_user_id"}))

	_, err := s.Gettaskbyid(999)
	if err == nil {
		t.Error("expected error for missing task ID, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
func TestStore_Getalltask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)

	}

	s := New(db)
	expectedTasks := []models.Task{
		{ID: 1, Name: "Test Task 1", Status: "pending", UserID: 10},
		{ID: 2, Name: "Test Task 2", Status: "completed", UserID: 20},
	}
	row := sqlmock.NewRows([]string{"taskid", "taskname", "status", "assigned_user_id"}).AddRow(expectedTasks[0].ID, expectedTasks[0].Name, expectedTasks[0].Status, expectedTasks[0].UserID).
		AddRow(expectedTasks[1].ID, expectedTasks[1].Name, expectedTasks[1].Status, expectedTasks[1].UserID)

	mock.ExpectQuery("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage").WillReturnRows(row)
	actual, err := s.Getalltask()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual[0].ID != expectedTasks[0].ID || actual[0].Name != expectedTasks[0].Name ||
		actual[0].Status != expectedTasks[0].Status || actual[0].UserID != expectedTasks[0].UserID || len(actual) != len(expectedTasks) {
		t.Errorf("Expected %+v, got %+v", expectedTasks, actual)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

}
func TestStore_Getalltask_ScanError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	rows := sqlmock.NewRows([]string{"taskid", "taskname", "status", "assigned_user_id"}).
		AddRow(1, nil, "pending", 10)

	mock.ExpectQuery("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage").
		WillReturnRows(rows)

	_, err := s.Getalltask()
	if err == nil {
		t.Error("expected scan error but got nil")
	}
}

func TestStore_Completetask(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)

	}
	s := New(db)
	taskid := 1
	mock.ExpectExec(`UPDATE taskmanage SET status = ? WHERE taskid = ?`).WithArgs("complete", taskid).WillReturnResult(sqlmock.NewResult(0, 1))

	msg, err := s.Completetask(taskid)
	if msg != "task completed successfully" {
		t.Errorf("expected success message, got: %s", msg)

	}
	if mock.ExpectationsWereMet() != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)

	}
}
func TestStore_Deletetask(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)

	}
	s := New(db)
	taskid := 1
	mock.ExpectExec("DELETE FROM taskmanage WHERE taskid = ?").WithArgs(taskid).WillReturnResult(sqlmock.NewResult(0, 1))
	msg, _ := s.Deletetask(taskid)
	if msg != "deleted task successfully" {
		t.Errorf("expected success message, got: %s", msg)
	}
}
func TestStore_Deletetask_NoRowAffected(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	mock.ExpectExec("DELETE FROM taskmanage").
		WithArgs(123).
		WillReturnResult(sqlmock.NewResult(0, 0))

	msg, err := s.Deletetask(123)
	if msg != "no task found" || err == nil {
		t.Errorf("expected no task found error, got: %s, err: %v", msg, err)
	}
}

func TestStore_Completetask_RowsAffectedError(t *testing.T) {
	db, mock, _ := sqlmock.New()

	s := New(db)

	mock.ExpectExec("UPDATE taskmanage SET status = \\? WHERE taskid = \\?").
		WithArgs("complete", 1).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("failed to get rows affected")))

	_, err := s.Completetask(1)
	if err == nil {
		t.Error("expected rows affected error, got nil")
	}
}
