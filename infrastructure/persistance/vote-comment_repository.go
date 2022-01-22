package persistance

import (
	"fmt"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type voteCommentRepository struct {
	sh *SqlHandler
}

func NewVoteCommentRepository(sh *SqlHandler) repository.VoteCommentRepository {
	return &voteCommentRepository{sh: sh}
}

// good/bad を増やす
func (vcR *voteCommentRepository) IncreaseVoteComment(vote *model.VoteComment) error {
	db := vcR.sh.db

	// user_id,comment_idで検索しcomment_votes tableにレコードが存在するか判定
	if err := db.Where("user_id = ? AND comment_id = ?", vote.UserID, vote.CommentID).First(&model.VoteComment{}).Error; err == nil {
		return fmt.Errorf("this uid already exists")
	}
	//以下のコードはなぜかうまくいかない...
	// if err:=db.First(&model.VoteComment{CommentID: vote.CommentID,UserID: vote.UserID}).Error;err!=nil{
	// 	return err
	// }

	// uidがあるかどうか
	if vote.UserID == "" {
		return fmt.Errorf("uid is empty")
	}
	//存在するか確認
	if err := db.First(&model.User{Id: vote.UserID}).Error; err != nil {
		return err
	}

	// DBからIDで検索し、commentにbind
	comment := &model.Comment{Id: vote.CommentID}
	if err := db.First(comment).Error; err != nil {
		return err
	}

	// VoteCntを更新
	if vote.IsUp {
		comment.VoteCnt = comment.VoteCnt + 1
	} else {
		comment.VoteCnt = comment.VoteCnt - 1
	}
	if err := db.Model(&model.Comment{Id: vote.CommentID}).Update(comment).Error; err != nil {
		return err
	}
	return db.Save(vote).Error
}

// good/bad の取り消し
func (vcR *voteCommentRepository) RevokeVoteComment(vote *model.VoteComment) error {
	db := vcR.sh.db

	// user_id,comment_idで検索しcomment_votes tableにレコードが存在するか判定
	if err := db.Where("user_id = ? AND comment_id = ?", vote.UserID, vote.CommentID).First(&model.VoteComment{}).Error; err != nil {
		return err
	}

	// DBからIDで検索し、commentにbind
	comment := &model.Comment{Id: vote.CommentID}
	if err := db.First(comment).Error; err != nil {
		return err
	}

	// VoteCntを更新
	if vote.IsUp {
		comment.VoteCnt = comment.VoteCnt - 1
	} else {
		comment.VoteCnt = comment.VoteCnt + 1
	}
	if err := db.Model(&model.Comment{Id: vote.CommentID}).Update(comment).Error; err != nil {
		return err
	}

	//下記のコードだと想定外のdeleteを行っている
	//おそらく 条件を user_id or commnet_idで削除してるのかな...
	// return db.Delete(vote).Error

	return db.Where("user_id = ? AND comment_id = ?", vote.UserID, vote.CommentID).Delete(&model.VoteComment{}).Error
}

// // good/bad済みか
// // もしすでにしてるものがあればそのcomment_idを返す
// func (vcR *voteCommentRepository) FindVoteCommentIdOfVoted(uid string,thread_id int) (*[]int,error){
// 	thread := &model.Thread{}

// }
