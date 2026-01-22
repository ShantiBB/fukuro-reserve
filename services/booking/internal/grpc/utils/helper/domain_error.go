package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"booking/internal/utils/consts"
)

var (
	errBookingNotFound      = status.Error(codes.NotFound, consts.ErrBookingNotFound.Error())
	errBookingRoomNotFound  = status.Error(codes.NotFound, consts.ErrBookingRoomNotFound.Error())
	errRoomLockAlreadyExist = status.Error(codes.AlreadyExists, consts.ErrRoomLockAlreadyExist.Error())
	errPriceChanged         = status.Error(codes.FailedPrecondition, consts.ErrPriceChanged.Error())
	errInternalServer       = status.Error(codes.Internal, consts.ErrInternalServer.Error())
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.ErrBookingNotFound):
		return errBookingNotFound

	case errors.Is(err, consts.ErrBookingRoomNotFound):
		return errBookingRoomNotFound

	case errors.Is(err, consts.ErrRoomLockAlreadyExist):
		return errRoomLockAlreadyExist

	case errors.Is(err, consts.ErrPriceChanged):
		return errPriceChanged

	default:
		return errInternalServer
	}
}
