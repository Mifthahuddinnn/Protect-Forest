package base

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  true,
		Data:    data,
	}
}
