package permission

import (
	"auth/internal/repository/models"
)

const (
	methodCreateUser = "/user.v1.UserService/CreateUser"
	methodGetUsers   = "/user.v1.UserService/GetUsers"
	methodGetUser    = "/user.v1.UserService/GetUser"
	methodUpdateUser = "/user.v1.UserService/UpdateUser"
	methodDeleteUser = "/user.v1.UserService/DeleteUser"
)

var (
	ownerOnly = map[string]struct{}{
		methodGetUser:    {},
		methodUpdateUser: {},
		methodDeleteUser: {},
	}
	moderatorOnly = map[string]struct{}{
		methodGetUsers: {},
		methodGetUser:  {},
	}
	adminOnly = map[string]struct{}{
		methodCreateUser: {},
		methodGetUsers:   {},
		methodGetUser:    {},
		methodUpdateUser: {},
		methodDeleteUser: {},
	}
)

func CheckRolePermission(method string, role models.UserRole, claimID, userID int64) bool {
	if role == models.UserRoleUser && claimID == userID {
		return checkOwnerMethods(method)
	}

	if role == models.UserRoleModerator {
		return checkModeratorMethods(method, claimID, userID)
	}

	if role == models.UserRoleAdmin {
		return checkAdminMethods(method)
	}

	return false
}

func checkOwnerMethods(method string) bool {
	_, ok := ownerOnly[method]
	return ok
}

func checkModeratorMethods(method string, claimID, userID int64) bool {
	if claimID == userID {
		if checkOwnerMethods(method) {
			return true
		}
	}

	_, ok := moderatorOnly[method]
	return ok
}

func checkAdminMethods(method string) bool {
	_, ok := adminOnly[method]
	return ok
}
