package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"booking/internal/utils/consts"
)

var (
	errInvalidHotelID   = status.Error(codes.InvalidArgument, consts.ErrInvalidHotelID.Error())
	errInvalidBookingID = status.Error(codes.InvalidArgument, consts.ErrInvalidBookingID.Error())
)
