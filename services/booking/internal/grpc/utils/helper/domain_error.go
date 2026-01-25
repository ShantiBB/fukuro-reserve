package helper

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"booking/internal/utils/consts"
)

type domainErr struct {
	message string
	code    codes.Code
}

var (
	errBookingNotFound      = domainErr{consts.MsgBookingNotFound, codes.NotFound}
	errBookingRoomNotFound  = domainErr{consts.MsgBookingRoomNotFound, codes.NotFound}
	errRoomLockAlreadyExist = domainErr{consts.MsgRoomLockAlreadyExist, codes.AlreadyExists}
	errPriceChanged         = domainErr{consts.MsgPriceChanged, codes.FailedPrecondition}
	errInternalServer       = domainErr{consts.MsgInternalServer, codes.Internal}
)

func HandleDomainErr(err error) error {
	if err == nil {
		return nil
	}

	var domErr domainErr
	switch {
	case errors.Is(err, consts.ErrBookingNotFound):
		domErr = errBookingNotFound
	case errors.Is(err, consts.ErrBookingRoomNotFound):
		domErr = errBookingRoomNotFound
	case errors.Is(err, consts.ErrRoomLockAlreadyExist):
		domErr = errRoomLockAlreadyExist
	case errors.Is(err, consts.ErrPriceChanged):
		domErr = errPriceChanged
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
