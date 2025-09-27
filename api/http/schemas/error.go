package schemas

type ValidateErrorResponse struct {
	Errors map[string]string `json:"errors"`
}
