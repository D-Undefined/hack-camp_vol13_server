package model

import (
	"time"
)

type Comment struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ThreadId  int       `json:"thread_id"`
	UserId    string    `json:"user_id"`
	Body      string    `json:"body"`
}
