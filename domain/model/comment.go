package model

import (
	"time"
)

type Comment struct {
	Id        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	ThreadID  int       `json:"thread_id"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
	Vote      int       `json:"vote"`
}
