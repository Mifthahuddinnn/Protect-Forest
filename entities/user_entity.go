package entities

import (
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password"form:"password"`
	Address   string `json:"address" form:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
