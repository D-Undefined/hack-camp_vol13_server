package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type VoteCommentRepository interface {
	IncreaseVoteComment(*model.VoteComment) error
	RevokeVoteComment(*model.VoteComment) error
}
