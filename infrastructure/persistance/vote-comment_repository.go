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

	// 投票した user
	vote_user := &model.User{Id: vote.UserID}
	//存在するか確認
	if err := db.First(vote_user).Error; err != nil {
		return err
	}

	// DBからIDで検索し、commentにbind
	comment := &model.Comment{Id: vote.CommentID}
	if err := db.First(comment).Error; err != nil {
		return err
	}

	// 投票されたコメントを作成した user
	wasVoted_user := &model.User{Id: comment.UserID}
	//存在するか確認
	if err := db.First(wasVoted_user).Error; err != nil {
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

	// userのスコアも更新
	vote_user.Score = vote_user.Score + 1
	if err := db.Model(&model.User{Id: vote.UserID}).Update(vote_user).Error; err != nil {
		return err
	}

	wasVoted_user.Score = wasVoted_user.Score + 2
	if err := db.Model(&model.User{Id: comment.UserID}).Update(wasVoted_user).Error; err != nil {
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

	// 投票した user
	vote_user := &model.User{Id: vote.UserID}
	//存在するか確認
	if err := db.First(vote_user).Error; err != nil {
		return err
	}

	// DBからIDで検索し、commentにbind
	comment := &model.Comment{Id: vote.CommentID}
	if err := db.First(comment).Error; err != nil {
		return err
	}

	// 投票されたコメントを作成した user
	wasVoted_user := &model.User{Id: comment.UserID}
	//存在するか確認
	if err := db.First(wasVoted_user).Error; err != nil {
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

	// userのスコアも更新 (元に戻す)
	vote_user.Score = vote_user.Score - 1
	if err := db.Model(&model.User{Id: vote.UserID}).Update(vote_user).Error; err != nil {
		return err
	}

	wasVoted_user.Score = wasVoted_user.Score - 2
	if err := db.Model(&model.User{Id: comment.UserID}).Update(wasVoted_user).Error; err != nil {
		return err
	}

	return db.Where("user_id = ? AND comment_id = ?", vote.UserID, vote.CommentID).Delete(&model.VoteComment{}).Error
}

// good/bad済みか
// もしすでにしてるものがあればそのcomment_idを返す
func (vcR *voteCommentRepository) FindVoteCommentIdOfVoted(uid string, threadId int) (*[]*model.VoteComment, error) {
	db := vcR.sh.db

	thread := &model.Thread{
		Comments: []*model.Comment{},
	}
	if err := db.Where(&model.Thread{Id: threadId}).Preload("Comments").Find(thread).Error; err != nil {
		return nil, err
	}

	// もっといい書き方があるかも
	thread_comment_id := []int{}
	for i := 0; i < len(thread.Comments); i++ {
		thread_comment_id = append(thread_comment_id, thread.Comments[i].Id)
	}

	vote_comments := &[]*model.VoteComment{}

	// 存在しなければ空の配列を返す
	if len(thread_comment_id) == 0 {
		return vote_comments, nil
	}

	if err := db.Where("user_id = ? AND comment_id IN (?) ", uid, thread_comment_id).Find(vote_comments).Error; err != nil {
		return nil, err
	}

	return vote_comments, nil

}
