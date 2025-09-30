package models

type UserCreate struct {
	Username    string `db:"username"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	Password    string `db:"password"`
}

type User struct {
	ID          int64  `db:"id"`
	Username    string `db:"username"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Description string `db:"description"`
}

type UserUpdatePassword struct {
	ID          int64
	Password    string
	NewPassword string
}

func (u UserCreate) ToUserRead() *User {
	return &User{
		Username:    u.Username,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Description: u.Description,
	}
}
