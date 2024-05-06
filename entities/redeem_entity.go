package entities

import "time"

type Redeem struct {
	ID             int  `json:"id" form:"id" gorm:"primaryKey"`
	UserID         int  `gorm:"not null"`
	User           User `gorm:"foreignKey:UserID"`
	PointsRedeemed int  `json:"points_redeemed" form:"points_redeemed"`
	AmountCredit   int  `json:"amount_credit" form:"amount_credit"`
	RedeemDate     time.Time
}
