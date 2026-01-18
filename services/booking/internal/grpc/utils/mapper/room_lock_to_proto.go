package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func RoomLockToProto(r *models.RoomLockShort) *bookingv1.RoomLockShort {
	lock := &bookingv1.RoomLockShort{
		Id:        r.ID.String(),
		IsActive:  r.ISActive,
		ExpiresAt: timestamppb.New(r.ExpiresAt),
		CreatedAt: timestamppb.New(r.CreatedAt),
	}

	return lock
}

func RoomLockWithDetailToProto(r *models.RoomLockDetail) *bookingv1.RoomLock {
	lock := &bookingv1.RoomLock{
		Id: r.ID.String(),
		StayRange: &bookingv1.DateRange{
			Start: timestamppb.New(r.StayRange.Start),
			End:   timestamppb.New(r.StayRange.End),
		},
		IsActive:  r.ISActive,
		ExpiresAt: timestamppb.New(r.ExpiresAt),
		CreatedAt: timestamppb.New(r.CreatedAt),
	}

	return lock
}
