package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int    `json:"id" form:"id" query:"id" gorm:"primaryKey"`
	FullName  string `json:"fullName" form:"fullName" query:"fullName"`
	Email     string `json:"email" form:"email" query:"email" gorm:"unique"`
	Password  string `json:"password" form:"password" query:"password"`
	Address   string `json:"address" form:"address" query:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
