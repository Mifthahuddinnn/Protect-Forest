package base

type LoginResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}

type TokenData struct {
	Token string `json:"token"`
}

func NewLoginResponse(message string, token string) *LoginResponse {
	return &LoginResponse{
		Message: message,
		Status:  true,
		Data:    TokenData{Token: token},
	}
}
