package entities

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name" `
	Email     string    `json:"email" form:"email" query:"email" gorm:"unique"`
	Password  string    `json:"password" form:"password" query:"password"`
	Address   string    `json:"address" form:"address" query:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time
}
