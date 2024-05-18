package response

import "forest/entities"

type UserResponse struct {
	ID        int    `json:"id" `
	Name      string `json:"name" `
	Email     string `json:"email"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

func FromUseCase(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
