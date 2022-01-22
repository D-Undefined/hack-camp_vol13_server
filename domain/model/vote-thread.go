package model

type VoteThread struct {
	Id       int    `json:"id" gorm:"primaryKey;not null"`
	UserID   string `gorm:"not null" json:"uid"`
	ThreadID int    `json:"thread_id"`
	IsUp     bool   `json:"is_up"`
}
