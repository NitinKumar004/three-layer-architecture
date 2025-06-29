package user

import (
	"errors"
	models_user "microservice/Models/user"
)

type Store interface {
	InsertUser(u models_user.User) (string, error)
	GetUserByID(id int) (*models_user.User, error)
	GetAllUsers() ([]models_user.User, error)
	DeleteAllUsers() (string, error)
	DeleteUserByID(id int) (string, error)
}

type service struct {
	store Store
}

func New(s Store) *service {
	return &service{store: s}
}

func (s *service) InsertUser(u models_user.User) (string, error) {
	if u.UserName == "" || u.UserPhone == "" || u.UserEmail == "" {
		return "", errors.New("all fields (name, phone, email) are required")
	}
	return s.store.InsertUser(u)
}

func (s *service) GetUserByID(id int) (*models_user.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.store.GetUserByID(id)
}

func (s *service) GetAllUsers() ([]models_user.User, error) {
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
