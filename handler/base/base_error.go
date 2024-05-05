package base

type ErrResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func NewErrResponse(message string) *ErrResponse {
	return &ErrResponse{
		Message: message,
		Status:  false,
	}
}
