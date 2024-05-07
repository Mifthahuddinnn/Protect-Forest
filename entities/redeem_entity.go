package entities

import (
	"time"
)

type Redeem struct {
	ID             int
	UserID         int
	User           User
	PointsRedeemed int
	AmountCredit   int
	RedeemDate     time.Time
}
