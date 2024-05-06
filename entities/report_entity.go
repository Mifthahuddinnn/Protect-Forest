package entities

type Reporting struct {
	ID          int    `json:"id" form:"id" gorm:"primaryKey"`
	UserID      int    `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	AdminID     int    `gorm:"not null"`
	Admin       Admin  `gorm:"foreignKey:AdminID"`
	Status      string `json:"status" form:"status"`
	Photo       string `json:"photo" form:"photo"`
	Description string `json:"description" form:"description"`
}
