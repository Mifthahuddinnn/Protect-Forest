package entities

import (
	"time"
)

type User struct {
	ID        int      `gorm:"primaryKey"`
	Name      string   `json:"name" form:"name"`
	Email     string   `json:"email" form:"email"`
	Password  string   `json:"password" form:"password"`
	Address   string   `json:"address" form:"address"`
	Points    int      `default:"0" gorm:"default:0"`
	Reports   []Report `gorm:"foreignKey:UserID"`
	Redeems   []Redeem `gorm:"foreignKey:UserID"`
	Balance   Balance  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
