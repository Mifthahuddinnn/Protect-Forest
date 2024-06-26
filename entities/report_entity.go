package entities

import "time"

type Report struct {
	ID            int       `gorm:"primaryKey"`
	UserID        int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User          User      `gorm:"foreignKey:UserID;references:ID"`
	AdminID       *int      `json:"admin_id,omitempty" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title         string    `json:"title" form:"title"`
	Content       string    `json:"content" form:"content"`
	ForestAddress string    `json:"forest" form:"forest"`
	Description   string    `json:"description" form:"description"`
	Photo         string    `json:"photo" form:"photo"`
	Status        string    `json:"status" form:"status" default:"pending" gorm:"default:pending"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time
}
