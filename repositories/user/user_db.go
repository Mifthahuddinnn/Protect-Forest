package user

import "time"

type UserDB struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name" `
	Email     string    `json:"email" form:"email" gorm:"unique"`
	Password  string    `json:"password" form:"password"`
	Address   string    `json:"address" form:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time
}
