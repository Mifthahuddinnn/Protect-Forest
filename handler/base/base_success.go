package base

type SuccessResponse struct {
	Message string
	Status  bool
	Data    interface{}
}

func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  true,
		Data:    data,
	}
}
