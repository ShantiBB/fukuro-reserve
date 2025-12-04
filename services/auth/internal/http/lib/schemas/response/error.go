package response

type ErrorSchema struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func ErrorResp(err error) *ErrorSchema {
	return &ErrorSchema{Type: "error", Message: err.Error()}
}
