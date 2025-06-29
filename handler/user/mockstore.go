package user

import (
	usermodel "microservice/Models/user"
)

type MockStore struct {
	Insertnewuserfunc  func(usermodel.User) (string, error)
	GetUserByIDfunc    func(id int) (*usermodel.User, error)
	Getalluserfunc     func() ([]usermodel.User, error)
	Deletealluserfunc  func() (string, error)
	DeleteUserbyidfunc func(id int) (string, error)
}

func (m *MockStore) InsertUser(u usermodel.User) (string, error) {
	return m.Insertnewuserfunc(u)

}
func (m *MockStore) GetUserByID(id int) (*usermodel.User, error) {
	return m.GetUserByIDfunc(id)

}

func (m *MockStore) GetAllUsers() ([]usermodel.User, error) {
	return m.Getalluserfunc()
}

func (m *MockStore) DeleteAllUsers() (string, error) {
	return m.Deletealluserfunc()
}

func (m *MockStore) DeleteUserByID(id int) (string, error) {
	return m.DeleteUserbyidfunc(id)
}
