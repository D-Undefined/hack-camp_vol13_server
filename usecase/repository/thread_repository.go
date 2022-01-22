package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type ThreadRepository interface {
	CreateThread(*model.Thread) error
	DeleteThread(*model.Thread) error
	UpdateThread(*model.Thread) error
	FindThreadById(int) (*model.Thread, error)
	FindAllThread() (*[]*model.Thread, error)
	FindTrendThread() (*[]*model.Thread, error)
	// UserOfThreadRanking() (*[]*model.UserRanking, error)
}
