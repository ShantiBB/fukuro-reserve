package helper

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"auth/pkg/lib/utils/consts"
)

type domainErr struct {
	message string
	code    codes.Code
}

var (
	errUserNotFound        = domainErr{consts.MsgUserNotFound, codes.NotFound}
	errUniqueUserField     = domainErr{consts.MsgUniqueUserField, codes.AlreadyExists}
	errInvalidID           = domainErr{consts.MsgInvalidID, codes.InvalidArgument}
	errPasswordHashing     = domainErr{consts.MsgPasswordHashing, codes.InvalidArgument}
	errInvalidCredentials  = domainErr{consts.MsgInvalidCredentials, codes.PermissionDenied}
	errInvalidBearer       = domainErr{consts.MsgInvalidBearer, codes.Unauthenticated}
	errInvalidRefreshToken = domainErr{consts.MsgInvalidToken, codes.Unauthenticated}
	errInvalidUnauthorized = domainErr{consts.MsgUnauthorized, codes.Unauthenticated}
	errForbidden           = domainErr{consts.MsgForbidden, codes.PermissionDenied}
	errInternalServer      = domainErr{consts.MsgInternalServer, codes.Internal}
)

func HandleDomainErr(err error) error {
	if err == nil {
		return nil
	}

	var domErr domainErr
	switch {
	case errors.Is(err, consts.ErrUserNotFound):
		domErr = errUserNotFound
	case errors.Is(err, consts.ErrUniqueUserField):
		domErr = errUniqueUserField
	case errors.Is(err, consts.ErrPasswordHashing):
		domErr = errPasswordHashing
	case errors.Is(err, consts.ErrInvalidID):
		domErr = errInvalidID
	case errors.Is(err, consts.ErrInvalidCredentials):
		domErr = errInvalidCredentials
	case errors.Is(err, consts.ErrInvalidBearer):
		domErr = errInvalidBearer
	case errors.Is(err, consts.ErrInvalidToken):
		domErr = errInvalidRefreshToken
	case errors.Is(err, consts.ErrUnauthorized):
		domErr = errInvalidUnauthorized
	case errors.Is(err, consts.ErrForbidden):
		domErr = errForbidden

	default:
		domErr = errInternalServer
	}

	ei := &errdetails.ErrorInfo{
		Reason: domErr.message,
		Domain: "user-service",
	}

	st, _ := status.New(domErr.code, "operation failed").WithDetails(ei)
	return st.Err()
}
