package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type VoteCommentRepository interface {
	IncreaseCommentVote(*model.CommentVote) error
	RevokeCommentVote(*model.CommentVote) error
}
