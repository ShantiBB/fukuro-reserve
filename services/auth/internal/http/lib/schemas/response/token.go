package response

type Token struct {
	Access    string `json:"access_token"`
	Refresh   string `json:"refresh_token"`
	TokenType string `json:"token_type" example:"Bearer"`
}
