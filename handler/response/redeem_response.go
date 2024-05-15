package response

import "forest/entities"

type RedeemResponse struct {
	ID    int
	Name  string
	Email string
}

func NewRedeemResponse(user *entities.User) RedeemResponse {
	return RedeemResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
