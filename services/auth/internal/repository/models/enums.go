package models

type UserRole string

const (
	UserRoleUnspecified UserRole = "USER_ROLE_UNSPECIFIED"
	UserRoleUser        UserRole = "USER_ROLE_USER"
	UserRoleModerator   UserRole = "USER_ROLE_MODERATOR"
	UserRoleAdmin       UserRole = "USER_ROLE_ADMIN"
)
