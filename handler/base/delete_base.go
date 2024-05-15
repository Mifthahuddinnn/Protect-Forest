package base

type DeleteBase struct {
	Message string
	Status  bool
}

func NewDeleteBase(message string) *DeleteBase {
	return &DeleteBase{
		Message: message,
		Status:  true,
	}
}
