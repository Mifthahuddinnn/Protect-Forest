package entities

type Admin struct {
	ID       int    `json:"id" form:"id" gorm:"primaryKey"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
