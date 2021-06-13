package model

import "time"

type Request struct {
	ID             int32     `gorm:"primaryKey" json:"id"`
	Message        string    `json:"message"`
	Intent         string    `json:"intent"`
	SenderId       string    `json:"sender_id"`
	ResponseAction string    `json:"response_action"`
	CreatedAt      time.Time `json:"created_at"`
}

func (Request) TableName() string {
	return "request"
}

const (
	RequestTypeAll          = 0
	RequestTypeClassified   = 1
	RequestTypeUnClassified = 2
)
