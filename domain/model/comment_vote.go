package model

type CommentVote struct {
	Id        int    `json:"id" gorm:"primaryKey;not null"`
	UserID    string `gorm:"not null" json:"uid"`
	CommentID int    `json:"comment_id"`
	IsUp      bool   `json:"is_up"`
}
