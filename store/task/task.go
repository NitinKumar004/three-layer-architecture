package task

import (
	"database/sql"
	"errors"
	"fmt"
	models_task "microservice/Models/task"
)

// This is defining a custom structure called Store.
// It has one field: db, which is a pointer to a SQL database connection.
// You don’t want to pass *sql.DB all over your code —
// instead, you wrap it in Store and put your DB-related functions there (like InsertTask, GetTaskByID etc.)
type Store struct {
	db *sql.DB
}

// This is a constructor function — it's a function that builds and returns a new Store object.

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Insertask(t models_task.Task) (string, error) {
	fmt.Println(t)
	_, err := s.db.Exec("INSERT INTO taskmanage(taskid, taskname, status, assigned_user_id)VALUES(?,?,?,?)", t.TaskID, t.TaskName, t.TaskStatus, t.AssignUser)
	if err != nil {
		return "data not inserted may be primary key issue here is issue : ", err
	}
	return "data inserted successfully", nil
}

func (s *Store) Gettaskbyid(id int) (*models_task.Task, error) {
	var t models_task.Task
	row := s.db.QueryRow("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage where taskid=?", id)

	err := row.Scan(&t.TaskID, &t.TaskName, &t.TaskStatus, &t.AssignUser)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Store) Getalltask() ([]models_task.Task, error) {
	var alltask []models_task.Task
	rowdata, err := s.db.Query("SELECT taskid, taskname, status, assigned_user_id FROM taskmanage")
	if err != nil {
		return nil, errors.New("fetching to all data")
	}
	for rowdata.Next() {
		var t models_task.Task
		err := rowdata.Scan(&t.TaskID, &t.TaskName, &t.TaskStatus, &t.AssignUser)
		if err != nil {
			return nil, errors.New("erro to fetching all data")
		}
		alltask = append(alltask, t)
	}
	return alltask, nil

}

func (s *Store) Deletetask(id int) (string, error) {

	result, err := s.db.Exec("DELETE FROM taskmanage WHERE taskid = ?", id)
	if err != nil {
		return "error to delete task", errors.New("failed to delete task")
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check update result")
	}
	if rowsaffected == 0 {
		return "no task found", errors.New("no task with this ID exists")
	}
	return "deleted task successfully", nil
}

func (s *Store) Completetask(id int) (string, error) {
	fmt.Println("id is ", id)
	result, err := s.db.Exec("UPDATE taskmanage SET status = ? WHERE taskid = ?", "complete", id)
	if err != nil {
		return "error to complete task", errors.New("failed to complete task")
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check update result")
	}
	if rowsaffected == 0 {
		return "no task found", errors.New("no task with this ID exists")
	}

	return "task completed successfully", nil
}
