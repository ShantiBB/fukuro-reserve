package constants

const (
	ErrCountryCodeRequired = "country_code is required"
	ErrCitySlugRequired    = "city_slug is required"
	ErrTitleRequired       = "title is required"
	ErrAddressRequired     = "address is required"
	ErrOwnerIDPositive     = "owner_id must be greater than 0"
	ErrLatitudeOutOfRange  = "location.latitude must be between -90 and 90"
	ErrLongitudeOutOfRange = "location.longitude must be between -180 and 180"
	ErrHotelUpdateRequired = "at least one of description, address or location is required"
)
