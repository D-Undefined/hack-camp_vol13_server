package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type VoteThreadRepository interface {
	IncreaseVoteThread(*model.VoteThread) error
	RevokeVoteThread(*model.VoteThread) error
	CheckVoteThread(string, int) (*model.VoteThread, error)
}
