package entities

import "time"

type Redeem struct {
	ID         int `json:"id" gorm:"primaryKey"`
	UserID     int `json:"user_id" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RedeemDate time.Time
}
