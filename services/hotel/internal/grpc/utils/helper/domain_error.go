package helper

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hotel/pkg/lib/utils/consts"
)

type domainErr struct {
	message string
	code    codes.Code
}

var (
	errHotelNotFound    = domainErr{consts.MsgHotelNotFound, codes.NotFound}
	errRoomNotFound     = domainErr{consts.MsgRoomNotFound, codes.NotFound}
	errUniqueHotelField = domainErr{consts.MsgUniqueHotelField, codes.NotFound}
	errUniqueRoomField  = domainErr{consts.MsgUniqueRoomField, codes.NotFound}
	errInternalServer   = domainErr{consts.MsgInternalServer, codes.Internal}
)

func HandleDomainErr(err error) error {
	if err == nil {
		return nil
	}

	var domErr domainErr
	switch {
	case errors.Is(err, consts.ErrHotelNotFound):
		domErr = errHotelNotFound
	case errors.Is(err, consts.ErrRoomNotFound):
		domErr = errRoomNotFound
	case errors.Is(err, consts.ErrUniqueHotelField):
		domErr = errUniqueHotelField
	case errors.Is(err, consts.ErrUniqueRoomField):
		domErr = errUniqueRoomField

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
