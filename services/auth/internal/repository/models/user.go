package models

import "time"

type CreateUser struct {
	Username *string
	Email    string
	Password string
}

type UpdateUser struct {
	Username string
	Email    string
	ID       int64
}

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  *string
	Email     string
	Role      UserRole
	ID        int64
	IsActive  bool
}

type UserShort struct {
	Username *string
	Email    string
	Role     UserRole
	ID       int64
	IsActive bool
}

type UserList struct {
	Users      []*UserShort
	TotalCount uint64
}

type UpdateUserPassword struct {
	Password    string
	NewPassword string
	ID          int64
}

type UserCredentials struct {
	Email    string
	Role     UserRole
	Password string
	ID       int64
}

func (u CreateUser) ToUserRead() *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
	}
}
