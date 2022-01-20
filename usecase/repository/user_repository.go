package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type UserRepository interface {
	FindAllUser() ([]*model.User, error)
	CreateUser(*model.User) error
	FindUserById(string) (*model.User, error)
}
