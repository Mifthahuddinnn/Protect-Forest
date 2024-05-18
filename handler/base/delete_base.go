package base

type DeleteBase struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func NewDeleteBase(message string) *DeleteBase {
	return &DeleteBase{
		Message: message,
		Status:  true,
	}
}
