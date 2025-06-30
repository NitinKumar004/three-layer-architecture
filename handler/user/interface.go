package user

import models_user "microservice/Models/user"

type service interface {
	InsertUser(u models_user.User) (string, error)
	GetUserByID(id int) (*models_user.User, error)
	GetAllUsers() ([]models_user.User, error)
	DeleteAllUsers() (string, error)
	DeleteUserByID(id int) (string, error)
}
