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
	if comment.UserID == "" {
		return fmt.Errorf("uid is empty")
	}

	user := &model.User{Id: comment.UserID}

	//userが存在するか確認
	if err := db.First(user).Error; err != nil {
		return err
	}

	//threadが存在するか確認
	thread := &model.Thread{Id: comment.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// commentCntを1増やす
	thread.CommentCnt = thread.CommentCnt + 1
	if err := db.Model(&model.Thread{Id: comment.ThreadID}).Update(thread).Error; err != nil {
		return err
	}

	// 投稿 userのScoreを5増やす
	user.Score = user.Score + 5
	if err := db.Model(&model.User{Id: comment.UserID}).Update(user).Error; err != nil {
		return err
	}

	return db.Save(comment).Error
}

func (cR commentRepository) DeleteComment(comment *model.Comment) error {
	db := cR.sh.db

	//threadが存在するか確認
	thread := &model.Thread{Id: comment.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// commentCntを1減らす
	thread.CommentCnt = thread.CommentCnt - 1
	if err := db.Model(&model.Thread{Id: comment.ThreadID}).Update(thread).Error; err != nil {
		return err
	}

	return db.Delete(comment).Error
}
