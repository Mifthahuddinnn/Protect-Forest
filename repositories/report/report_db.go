package report

import (
	"forest/repositories/admin"
	"forest/repositories/user"
	"time"
)

type ReportingDB struct {
	ID          int         `json:"id" form:"id" gorm:"primaryKey"`
	UserID      int         `gorm:"null"`
	User        user.UserDB `gorm:"foreignKey:UserID"`
	AdminID     int         `gorm:"null"`
	Admin       admin.DB    `gorm:"foreignKey:AdminID"`
	Status      string      `json:"status" form:"status"`
	Photo       string      `json:"photo" form:"photo"`
	Description string      `json:"description" form:"description"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time
}
