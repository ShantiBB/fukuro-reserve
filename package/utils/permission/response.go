package permission

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorResp(msg string) *ErrorResponse {
	return &ErrorResponse{Error: msg}
}
