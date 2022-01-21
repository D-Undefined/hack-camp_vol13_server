package model

import (
	"time"
)

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	ThreadID  int       `json:"thread_id"`
	UserID    string    `json:"uid"`
	Body      string    `json:"body"`
	Vote      int       `json:"vote"`
}
