package schemas

type UserCreateRequest struct {
	Username    string `json:"username" validate:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" validate:"required,email"`
	Description string `json:"description"`
	Password    string `json:"password" validate:"required,min=5"`
}

type UserRequest struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" validate:"email"`
	Description string `json:"description"`
}
