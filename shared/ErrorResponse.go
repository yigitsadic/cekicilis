package shared

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}
