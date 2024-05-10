package entities

import "time"

type Admin struct {
	ID        int
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
