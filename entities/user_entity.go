package entities

import (
	"time"
)

type User struct {
	ID          int    `json:"id" form:"id" gorm:"primaryKey"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	Points      int    `json:"points" form:"points" default:"0"`
	Balance     int    `json:"balance" form:"balance" default:"0"`
	Reports     []Reporting
	Redemptions []Redeem
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
