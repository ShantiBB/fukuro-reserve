package validation

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

var (
	currencyRegexp = regexp.MustCompile(`^[A-Z]{3}$`)
	amountRegexp   = regexp.MustCompile(`^[0-9]+(\.[0-9]{1,18})?$`)
)

// ValidateUUID validates that a string is a valid UUID format.
func ValidateUUID(value string) error {
	if _, err := uuid.Parse(value); err != nil {
		return errors.New("must be a valid UUID format")
	}
	return nil
}

// ValidateCurrency validates that a string matches the currency pattern (3 uppercase letters).
func ValidateCurrency(value string) error {
	if !currencyRegexp.MatchString(value) {
		return errors.New("must be a valid 3-letter currency code (e.g., USD, EUR)")
	}
	return nil
}

// ValidateAmount validates that a string matches the amount pattern (numeric with up to 18 decimal places).
func ValidateAmount(value string) error {
	if !amountRegexp.MatchString(value) {
		return errors.New("must be a valid numeric amount (up to 18 decimal places)")
	}
	return nil
}

func ValidateRegisterRequest(req dto.RegisterRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf("email %w", err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}

func ValidateLoginRequest(req dto.LoginRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf("email %w", err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}

func ValidateRefreshTokenRequest(req dto.RefreshTokenRequest) error {
	if strings.TrimSpace(req.RefreshToken) == "" {
		return errors.New("refresh_token is required")
	}
	return nil
}

func ValidateCreateUserRequest(req dto.CreateUserRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf("email %w", err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}

func ValidateUpdateUserRequest(req dto.UpdateUserRequest) error {
	if strings.TrimSpace(req.Email) == "" && strings.TrimSpace(req.Username) == "" {
		return errors.New("at least one of email or username is required")
	}
	if strings.TrimSpace(req.Email) != "" {
		if err := validateEmail(req.Email); err != nil {
			return fmt.Errorf("email %w", err)
		}
	}
	return nil
}

func ValidateUpdateUserRoleRequest(req dto.UpdateUserRoleRequest) error {
	if strings.TrimSpace(req.Role) == "" {
		return errors.New("role is required")
	}
	return nil
}

func ValidateCreateHotelRequest(req dto.CreateHotelRequest) error {
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New("country_code is required")
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New("city_slug is required")
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(req.Address) == "" {
		return errors.New("address is required")
	}
	if req.OwnerId <= 0 {
		return errors.New("owner_id must be greater than 0")
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New("location.latitude must be between -90 and 90")
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New("location.longitude must be between -180 and 180")
		}
	}
	return nil
}

func ValidateUpdateHotelRequest(req dto.UpdateHotelRequest) error {
	if strings.TrimSpace(req.Description) == "" && strings.TrimSpace(req.Address) == "" && req.Location == nil {
		return errors.New("at least one of description, address or location is required")
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New("location.latitude must be between -90 and 90")
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New("location.longitude must be between -180 and 180")
		}
	}
	return nil
}

func ValidateUpdateHotelTitleRequest(req dto.UpdateHotelTitleRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title is required")
	}
	return nil
}

func ValidateCreateRoomRequest(req dto.CreateRoomRequest) error {
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New("room_number is required")
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New("type is required")
	}
	if strings.TrimSpace(req.HotelSlug) == "" {
		return errors.New("hotel_slug is required")
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New("country_code is required")
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New("city_slug is required")
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf("price %w", err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New("capacity must be between 1 and 20")
	}
	if req.Floor < 0 {
		return errors.New("floor must be greater than or equal to 0")
	}
	if req.AreaSqm <= 0 {
		return errors.New("area_sqm must be greater than 0")
	}
	return nil
}

func ValidateUpdateRoomRequest(req dto.UpdateRoomRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New("room_number is required")
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New("type is required")
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf("price %w", err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New("capacity must be between 1 and 20")
	}
	if req.Floor < 0 {
		return errors.New("floor must be greater than or equal to 0")
	}
	if req.AreaSqm <= 0 {
		return errors.New("area_sqm must be greater than 0")
	}
	return nil
}

func ValidateUpdateRoomStatusRequest(req dto.UpdateRoomStatusRequest) error {
	if strings.TrimSpace(req.Status) == "" {
		return errors.New("status is required")
	}
	return nil
}

func ValidateCreateBookingRequest(req dto.CreateBookingRequest) error {
	if strings.TrimSpace(req.GuestName) == "" {
		return errors.New("guest_name is required")
	}
	if err := ValidateUUID(req.HotelId); err != nil {
		return fmt.Errorf("hotel_id %w", err)
	}
	if err := ValidateCurrency(req.Currency); err != nil {
		return fmt.Errorf("currency %w", err)
	}
	if err := ValidateAmount(req.ExpectedTotalAmount); err != nil {
		return fmt.Errorf("expected_total_amount %w", err)
	}
	if !req.CheckIn.After(time.Now()) {
		return errors.New("check_in must be in the future")
	}
	if !req.CheckOut.After(req.CheckIn) {
		return errors.New("check_out must be after check_in")
	}
	if len(req.Rooms) == 0 {
		return errors.New("at least one room is required")
	}

	for i, room := range req.Rooms {
		if room == nil {
			return fmt.Errorf("rooms[%d] is required", i)
		}
		if err := ValidateRoom(room); err != nil {
			return fmt.Errorf("rooms[%d]: %w", i, err)
		}
	}

	if req.GuestEmail != "" {
		if err := validateEmail(req.GuestEmail); err != nil {
			return fmt.Errorf("guest_email %w", err)
		}
	}

	return nil
}

func ValidateRoom(room *dto.CreateBookingRoomRequest) error {
	if err := ValidateUUID(room.RoomId); err != nil {
		return fmt.Errorf("room_id %w", err)
	}
	if room.Adults < 1 || room.Adults > 20 {
		return errors.New("adults must be between 1 and 20")
	}
	if room.Children > 20 {
		return errors.New("children must be 20 or less")
	}
	if err := ValidateAmount(room.PricePerNight); err != nil {
		return fmt.Errorf("price_per_night %w", err)
	}

	return nil
}

func validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("must be a valid email")
	}
	return nil
}
