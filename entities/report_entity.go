package entities

import "time"

type Reporting struct {
	ID          int
	UserID      int
	User        User
	AdminID     int
	Admin       Admin
	Status      string
	Photo       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
