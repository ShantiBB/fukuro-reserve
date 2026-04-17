package validation

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

// ValidateUUID validates that a string is a valid UUID format
func ValidateUUID(value string) error {
	if _, err := uuid.Parse(value); err != nil {
		return errors.New("must be a valid UUID format")
	}
	return nil
}

// ValidateCurrency validates that a string matches the currency pattern (3 uppercase letters)
func ValidateCurrency(value string) error {
	matched, _ := regexp.MatchString(`^[A-Z]{3}$`, value)
	if !matched {
		return errors.New("must be a valid 3-letter currency code (e.g., USD, EUR)")
	}
	return nil
}

// ValidateAmount validates that a string matches the amount pattern (numeric with up to 18 decimal places)
func ValidateAmount(value string) error {
	matched, _ := regexp.MatchString(`^[0-9]+(\.[0-9]{1,18})?$`, value)
	if !matched {
		return errors.New("must be a valid numeric amount (up to 18 decimal places)")
	}
	return nil
}

// CreateBookingRoomRequest represents a room in a booking request
// This is a copy of the type from handler package to avoid import cycles
type CreateBookingRoomRequest struct {
	RoomId        string
	PricePerNight string
	Adults        uint32
	Children      uint32
}

// ValidateRoom validates a single room's fields
func ValidateRoom(room *CreateBookingRoomRequest) error {
	if room.RoomId != "" {
		if err := ValidateUUID(room.RoomId); err != nil {
			return err
		}
	}

	if room.Adults < 1 || room.Adults > 20 {
		return errors.New("adults must be between 1 and 20")
	}

	if room.Children > 20 {
		return errors.New("children must be 20 or less")
	}

	if room.PricePerNight != "" {
		if err := ValidateAmount(room.PricePerNight); err != nil {
			return err
		}
	}

	return nil
}
