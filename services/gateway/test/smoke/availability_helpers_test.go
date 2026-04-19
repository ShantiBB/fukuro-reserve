package smoke

import (
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func findRoomByID(rooms []*dto.RoomShortResponse, id string) *dto.RoomShortResponse {
	for _, room := range rooms {
		if room != nil && room.Id == id {
			return room
		}
	}
	return nil
}

func availabilityPath(env *fixtures.Env, checkIn, checkOut time.Time) string {
	return availabilityBasePath(env) +
		"?check_in=" + checkIn.Format("2006-01-02") +
		"&check_out=" + checkOut.Format("2006-01-02") +
		"&page=1&limit=10"
}

func wrongLocationAvailabilityPath(env *fixtures.Env, checkIn, checkOut time.Time) string {
	return wrongLocationAvailabilityBasePath(env) +
		"?check_in=" + checkIn.Format("2006-01-02") +
		"&check_out=" + checkOut.Format("2006-01-02") +
		"&page=1&limit=10"
}
