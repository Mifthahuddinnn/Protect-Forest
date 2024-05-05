package response

import "forest/entities"

type UserResponse struct {
	ID        int    `json:"id" form:"id" gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Address   string `json:"address" form:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromUseCase(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
