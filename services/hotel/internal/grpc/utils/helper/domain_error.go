package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hotel/pkg/lib/utils/consts"
)

var (
	errBookingNotFound = status.Error(codes.NotFound, consts.MsgHotelNotFound)
	errInternalServer  = status.Error(codes.Internal, consts.MsgInternalServer)
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.ErrHotelNotFound):
		return errBookingNotFound

	default:
		return errInternalServer
	}
}
