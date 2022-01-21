package persistance

import (
	"fmt"
	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type commentRepository struct {
	sh *SqlHandler
}

func NewCommentRepository(sh *SqlHandler) repository.CommentRepository {
	return &commentRepository{sh: sh}
}

func (cR commentRepository) CreateComment(comment *model.Comment) error {
	db := cR.sh.db

	// uidがあるかどうか
	if comment.UserID==""{
		return fmt.Errorf("uid is empty")
	}
	//userが存在するか確認
	if err := db.First(&model.User{Id: comment.UserID}).Error; err != nil {
		return err
	}

	//threadが存在するか確認
	if err := db.First(&model.Thread{Id: comment.ThreadID}).Error; err != nil {
		return err
	}
	return db.Save(comment).Error
}

func (cR commentRepository) DeleteComment(comment *model.Comment) error {
	db := cR.sh.db
	return db.Delete(comment).Error
}
