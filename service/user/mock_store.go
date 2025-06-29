package user

import (
	usermodel "microservice/Models/user"
)

type Mockstore struct {
	Insertnewuserfunc  func(usermodel.User) (string, error)
	GetUserByIDfunc    func(id int) (*usermodel.User, error)
	Getalluserfunc     func() ([]usermodel.User, error)
	Deletealluserfunc  func() (string, error)
	DeleteUserbyidfunc func(id int) (string, error)
}

func (m *Mockstore) InsertUser(u usermodel.User) (string, error) {
	return m.Insertnewuserfunc(u)

}
func (m *Mockstore) GetUserByID(id int) (*usermodel.User, error) {
	return m.GetUserByIDfunc(id)

}

func (m *Mockstore) GetAllUsers() ([]usermodel.User, error) {
	return m.Getalluserfunc()
}

func (m *Mockstore) DeleteAllUsers() (string, error) {
	return m.Deletealluserfunc()
}

func (m *Mockstore) DeleteUserByID(id int) (string, error) {
	return m.DeleteUserbyidfunc(id)
}
