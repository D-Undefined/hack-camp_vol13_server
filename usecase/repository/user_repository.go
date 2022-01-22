package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type UserRepository interface {
	CreateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
	UpdateUser(*model.User) error
	FindUserById(string) (*model.User, error)
	FindAllUser() (*[]*model.User, error)
	GetUserRanking() (*[]*model.User, error)
}
