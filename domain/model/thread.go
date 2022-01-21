package model

import (
	"time"
)

type Thread struct {
	Id         int       `json:"id" gorm:"primaryKey;not null"`
	Name       string    `json:"name"`
	VoteGood       int       `json:"vote_good"`
	VoteBad       int       `json:"vote_bad"`
	UserID     string    `gorm:"not null" json:"uid"`
	CreatedAt  time.Time `json:"created_at"`
	CommentCnt int       `json:"comment_cnt"`
	Comments   []*Comment
}
