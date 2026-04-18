package helper

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/utils/consts"
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
	errInvalidHotelID       = domainErr{consts.MsgInvalidHotelID, codes.InvalidArgument}
	errInvalidBookingID     = domainErr{consts.MsgInvalidBookingID, codes.InvalidArgument}
	errInvalidBookingRoomID = domainErr{consts.MsgInvalidBookingRoomID, codes.InvalidArgument}
	errInvalidPricePerNight = domainErr{consts.MsgInvalidPricePerNightID, codes.InvalidArgument}
	errInvalidExpectedTotal = domainErr{consts.MsgInvalidExpectedTotalAmountID, codes.InvalidArgument}
	errInvalidBookingStatus = domainErr{consts.MsgInvalidBookingStatus, codes.InvalidArgument}
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
	case errors.Is(err, consts.ErrInvalidHotelID):
		domErr = errInvalidHotelID
	case errors.Is(err, consts.ErrInvalidBookingID):
		domErr = errInvalidBookingID
	case errors.Is(err, consts.ErrInvalidBookingRoomID):
		domErr = errInvalidBookingRoomID
	case errors.Is(err, consts.ErrInvalidPricePerNightID):
		domErr = errInvalidPricePerNight
	case errors.Is(err, consts.ErrInvalidExpectedTotalAmountID):
		domErr = errInvalidExpectedTotal
	case errors.Is(err, consts.ErrInvalidBookingStatus):
		domErr = errInvalidBookingStatus
	default:
		domErr = errInternalServer
	}

	ei := &errdetails.ErrorInfo{
		Reason: domErr.message,
		Domain: "booking-service",
	}

	st, _ := status.New(domErr.code, "operation failed").WithDetails(ei)
	return st.Err()
}
