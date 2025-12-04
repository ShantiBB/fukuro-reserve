package models

import "time"

type UserCreate struct {
	Username *string
	Email    string
	Password string
}

type User struct {
	ID        int64
	Username  *string
	Email     string
	Role      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdatePassword struct {
	ID          int64
	Password    string
	NewPassword string
}

type UserCredentials struct {
	ID       int64
	Email    string
	Role     string
	Password string
}

func (u UserCreate) ToUserRead() User {
	return User{
		Username: u.Username,
		Email:    u.Email,
	}
}
