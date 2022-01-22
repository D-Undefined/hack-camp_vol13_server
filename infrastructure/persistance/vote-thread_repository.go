package persistance

import (
	"fmt"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type voteThreadRepository struct {
	sh *SqlHandler
}

func NewVoteThreadRepository(sh *SqlHandler) repository.VoteThreadRepository {
	return &voteThreadRepository{sh: sh}
}

// good/bad を増やす
func (vtR *voteThreadRepository) IncreaseVoteThread(vote *model.VoteThread) error {
	db := vtR.sh.db

	// user_id,thread_idで検索し
	// thread_votes tableにレコードが存在するか判定
	if err := db.Where("user_id = ? AND thread_id = ?", vote.UserID, vote.ThreadID).First(&model.VoteThread{}).Error; err == nil {
		return fmt.Errorf("this uid and thread_id already exists")
	}

	// uidがあるかどうか
	if vote.UserID == "" {
		return fmt.Errorf("uid is empty")
	}
	//userが存在するか確認
	if err := db.First(&model.User{Id: vote.UserID}).Error; err != nil {
		return err
	}

	// DBからIDで検索し、threadにbind
	thread := &model.Thread{Id: vote.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// VoteCntを更新
	if vote.IsUp {
		thread.VoteCnt = thread.VoteCnt + 1
	} else {
		thread.VoteCnt = thread.VoteCnt - 1
	}
	if err := db.Model(&model.Thread{Id: vote.ThreadID}).Update(thread).Error; err != nil {
		return err
	}
	return db.Save(vote).Error
}

// good/bad の取り消し
func (vtR *voteThreadRepository) RevokeVoteThread(vote *model.VoteThread) error {
	db := vtR.sh.db

	// user_id,thread_idで検索し
	// thread_votes tableにレコードが存在するか判定
	if err := db.Where("user_id = ? AND thread_id = ?", vote.UserID, vote.ThreadID).First(&model.VoteThread{}).Error; err != nil {
		return err
	}

	// DBからIDで検索し、threadにbind
	thread := &model.Thread{Id: vote.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// VoteCntを更新
	if vote.IsUp {
		thread.VoteCnt = thread.VoteCnt - 1
	} else {
		thread.VoteCnt = thread.VoteCnt + 1
	}
	if err := db.Model(&model.Thread{Id: vote.ThreadID}).Update(thread).Error; err != nil {
		return err
	}

	return db.Where("user_id = ? AND thread_id = ?", vote.UserID, vote.ThreadID).Delete(&model.VoteThread{}).Error
}
