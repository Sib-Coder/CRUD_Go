package service

import (
	"CRUD_Go/internal/app/model"
)

type Database interface {
	ExtractUser(name string) (model.User, error)
	AddUser(user model.User) (string, error)
	DeleteUser(user model.User) (string, error)
	UpdateUser(user model.User) (string, error)
	ExtractUsers() ([]model.User, error)
}
type Service struct {
	//database1 *db.Database
	db Database
}

func New(db Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Takeuser(name string) (model.User, error) {
	user, err := s.db.ExtractUser(name)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *Service) Takeusers() ([]model.User, error) {
	users, err := s.db.ExtractUsers()
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (s *Service) AddUser(u model.User) (string, error) {
	user, err := s.db.AddUser(u)
	if err != nil {
		return " ", err
	}
	return user, nil
}
func (s *Service) Deleteuser(u model.User) (string, error) {
	user, err := s.db.DeleteUser(u)
	if err != nil {
		return " ", err
	}
	return user, nil
}
func (s *Service) Updateuser(u model.User) (string, error) {
	user, err := s.db.UpdateUser(u)
	if err != nil {
		return " ", err
	}
	return user, nil
}
