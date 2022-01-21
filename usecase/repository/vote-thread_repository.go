package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type VoteThreadRepository interface {
	IncreaseThreadVote(*model.ThreadVote) error
	RevokeThreadVote(*model.ThreadVote) error
}
