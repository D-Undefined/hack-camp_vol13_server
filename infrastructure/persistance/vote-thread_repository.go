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

	// 投票した user
	vote_user := &model.User{Id: vote.UserID}
	//userが存在するか確認
	if err := db.First(vote_user).Error; err != nil {
		return err
	}

	// DBからIDで検索し、threadにbind
	thread := &model.Thread{Id: vote.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// 投票された threadを作成した user
	wasVoted_user := &model.User{Id: thread.UserID}
	//userが存在するか確認
	if err := db.First(wasVoted_user).Error; err != nil {
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

	// userのスコアも更新
	vote_user.Score = vote_user.Score + 2
	if err := db.Model(&model.User{Id: vote.UserID}).Update(vote_user).Error; err != nil {
		return err
	}

	wasVoted_user.Score = wasVoted_user.Score + 4
	if err := db.Model(&model.User{Id: thread.UserID}).Update(wasVoted_user).Error; err != nil {
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

	// 投票した user
	vote_user := &model.User{Id: vote.UserID}
	//userが存在するか確認
	if err := db.First(vote_user).Error; err != nil {
		return err
	}

	// DBからIDで検索し、threadにbind
	thread := &model.Thread{Id: vote.ThreadID}
	if err := db.First(thread).Error; err != nil {
		return err
	}

	// 投票された threadを作成した user
	wasVoted_user := &model.User{Id: thread.UserID}
	//userが存在するか確認
	if err := db.First(wasVoted_user).Error; err != nil {
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

	// userのスコアも更新 (元に戻す)
	vote_user.Score = vote_user.Score - 2
	if err := db.Model(&model.User{Id: vote.UserID}).Update(vote_user).Error; err != nil {
		return err
	}

	wasVoted_user.Score = wasVoted_user.Score - 4
	if err := db.Model(&model.User{Id: thread.UserID}).Update(wasVoted_user).Error; err != nil {
		return err
	}

	return db.Where("user_id = ? AND thread_id = ?", vote.UserID, vote.ThreadID).Delete(&model.VoteThread{}).Error
}

// good/bad済みか
func (vtR *voteThreadRepository) CheckVoteThread(uid string, threadId int) (*model.VoteThread, error) {
	db := vtR.sh.db

	vote_thread := &model.VoteThread{}
	if err := db.Where(&model.VoteThread{ThreadID: threadId, UserID: uid}).Find(vote_thread).Error; err != nil {
		return &model.VoteThread{}, err
	}
	return vote_thread, nil
}
