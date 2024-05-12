package entities

type Balance struct {
	ID     int `gorm:"primaryKey"`
	UserID int `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Amount int `default:"0" gorm:"default:0"`
}
