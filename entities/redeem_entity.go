package entities

import "time"

type Redeem struct {
	ID          int `json:"id" gorm:"primaryKey"`
	PointNeeded int `json:"point" form:"point" default:"0" gorm:"default:0"`
	UserID      int `json:"user_id" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RedeemDate  time.Time
}
