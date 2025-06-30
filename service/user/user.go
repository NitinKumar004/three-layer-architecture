package user

import (
	"errors"
	User_Model "microservice/Models/user"
)

type service struct {
	store Store
}

func New(s Store) *service {
	return &service{store: s}
}

func (s *service) InsertUser(u User_Model.User) (string, error) {
	if u.Name == "" || u.Phone == "" || u.Email == "" {
		return "", errors.New("all fields (name, phone, email) are required")
	}
	return s.store.InsertUser(u)
}

func (s *service) GetUserByID(id int) (*User_Model.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.store.GetUserByID(id)
}

func (s *service) GetAllUsers() ([]User_Model.User, error) {
	return s.store.GetAllUsers()
}

func (s *service) DeleteAllUsers() (string, error) {
	return s.store.DeleteAllUsers()
}

func (s *service) DeleteUserByID(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("invalid user ID")
	}
	return s.store.DeleteUserByID(id)
}
