package base

type ReportResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func NewReportResponse(message string, data interface{}) *ReportResponse {
	return &ReportResponse{
		Message: message,
		Status:  true,
		Data:    data,
	}
}
