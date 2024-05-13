package entities

import "time"

type Message struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `json:"user_id" form:"user_id"`
	Content   string `json:"content" form:"content"`
	Direction string `json:"direction" form:"direction"`
	CreatedAt time.Time
}
