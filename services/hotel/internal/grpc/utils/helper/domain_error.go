package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hotel/pkg/lib/utils/consts"
)

var (
	errBookingNotFound  = status.Error(codes.NotFound, consts.MsgHotelNotFound)
	errUniqueHotelField = status.Error(codes.NotFound, consts.MsgUniqueHotelField)
	errInternalServer   = status.Error(codes.Internal, consts.MsgInternalServer)
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.ErrHotelNotFound):
		return errBookingNotFound
	case errors.Is(err, consts.ErrUniqueHotelField):
		return errUniqueHotelField

	default:
		return errInternalServer
	}
}
