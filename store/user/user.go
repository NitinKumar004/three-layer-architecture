package user

import (
	"database/sql"
	"errors"
	models_user "microservice/Models/user"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) InsertUser(u models_user.User) (string, error) {
	//fmt.Println(u)
	_, err := s.db.Exec("INSERT INTO usermanage(userid, username, userphone, useremail) VALUES (?, ?, ?, ?)",
		u.UserID, u.UserName, u.UserPhone, u.UserEmail)

	if err != nil {
		return "user not inserted, maybe primary key issue", err
	}
	return "user inserted successfully", nil
}

func (s *Store) GetUserByID(id int) (*models_user.User, error) {
	var u models_user.User
	row := s.db.QueryRow("SELECT userid, username, userphone, useremail FROM usermanage WHERE userid = ?", id)

	err := row.Scan(&u.UserID, &u.UserName, &u.UserPhone, &u.UserEmail)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Store) GetAllUsers() ([]models_user.User, error) {
	var allUsers []models_user.User
	rows, err := s.db.Query("SELECT userid, username, userphone, useremail FROM usermanage")
	if err != nil {
		return nil, errors.New("error fetching users")
	}

	for rows.Next() {
		var u models_user.User
		err := rows.Scan(&u.UserID, &u.UserName, &u.UserPhone, &u.UserEmail)
		if err != nil {
			return nil, errors.New("error scanning user data")
		}
		allUsers = append(allUsers, u)
	}
	return allUsers, nil
}

func (s *Store) DeleteAllUsers() (string, error) {
	result, err := s.db.Exec("DELETE FROM usermanage")
	if err != nil {
		return "error deleting users", errors.New("failed to delete users")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check deletion result")
	}
	if rowsAffected == 0 {
		return "no users found", errors.New("no users to delete")
	}

	return "all users deleted successfully", nil
}
func (s *Store) DeleteUserByID(id int) (string, error) {
	result, err := s.db.Exec("DELETE FROM usermanage WHERE userid = ?", id)
	if err != nil {
		return "error deleting user", errors.New("failed to delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error reading result", errors.New("failed to check deletion result")
	}
	if rowsAffected == 0 {
		return "no user found", errors.New("no user with this ID exists")
	}

	return "user deleted successfully", nil
}
