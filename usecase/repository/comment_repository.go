package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type CommentRepository interface {
	CreateComment(*model.Comment) error
	DeleteComment(*model.Comment) error
}
