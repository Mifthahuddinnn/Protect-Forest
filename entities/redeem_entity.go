package entities

import "time"

type Redeem struct {
	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RedeemDate time.Time
}
