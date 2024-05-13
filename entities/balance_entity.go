package entities

type Balance struct {
	ID     int `gorm:"primaryKey"`
	UserID int `gorm:"uniqueIndex;not null"` 
	Amount int `default:"0" gorm:"default:0"`
}
