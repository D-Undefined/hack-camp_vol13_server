package model

import (
	"time"
)

type Thread struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Vote      int       `json:"vote"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Comment   []*Comment
}
