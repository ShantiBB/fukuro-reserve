package dto

import (
	"time"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

// RegisterRequest Auth DTOs
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// User DTOs
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserActivityRequest struct {
	IsActive bool `json:"is_active"`
}

type UpdateUserActivityResponse struct {
	IsActive bool `json:"is_active"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

type UpdateUserRoleResponse struct {
	Role string `json:"role"`
}

type UserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username  string    `json:"username,omitempty"`
	Role      string    `json:"role"`
	Id        int64     `json:"id"`
	IsActive  bool      `json:"is_active"`
}

type UsersResponse struct {
	Users []*UserResponse `json:"users"`
}

// Hotel DTOs
type LocationDTO struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CreateHotelRequest struct {
	Location    *LocationDTO `json:"location,omitempty"`
	CountryCode string       `json:"country_code"`
	CitySlug    string       `json:"city_slug"`
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Address     string       `json:"address"`
	OwnerId     int64        `json:"owner_id"`
}

type UpdateHotelRequest struct {
	Location    *LocationDTO `json:"location,omitempty"`
	Description string       `json:"description"`
	Address     string       `json:"address"`
}

type UpdateHotelTitleRequest struct {
	Title string `json:"title"`
}

type HotelResponse struct {
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Location    *LocationDTO `json:"location,omitempty"`
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	HotelSlug   string       `json:"hotel_slug,omitempty"`
	Description string       `json:"description"`
	Address     string       `json:"address"`
	OwnerId     int64        `json:"owner_id"`
	Rating      float32      `json:"rating,omitempty"`
}

type HotelShortResponse struct {
	Location  *LocationDTO `json:"location,omitempty"`
	Id        string       `json:"id"`
	Title     string       `json:"title"`
	HotelSlug string       `json:"hotel_slug"`
	Address   string       `json:"address"`
	OwnerId   int64        `json:"owner_id"`
	Rating    float32      `json:"rating,omitempty"`
}

type HotelsResponse struct {
	Hotels []*HotelShortResponse `json:"hotels"`
}

// Room DTOs
type CreateRoomRequest struct {
	RoomNumber  string   `json:"room_number"`
	Type        string   `json:"type"`
	HotelSlug   string   `json:"hotel_slug"`
	HotelId     string   `json:"hotel_id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Price       string   `json:"price"`
	CountryCode string   `json:"country_code"`
	CitySlug    string   `json:"city_slug"`
	Amenities   []string `json:"amenities"`
	Images      []string `json:"images"`
	Floor       int64    `json:"floor"`
	Capacity    int64    `json:"capacity"`
	AreaSqm     float32  `json:"area_sqm"`
}

type UpdateRoomRequest struct {
	Title       string   `json:"title"`
	RoomNumber  string   `json:"room_number"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Amenities   []string `json:"amenities"`
	Images      []string `json:"images"`
	Capacity    int64    `json:"capacity"`
	Floor       int64    `json:"floor"`
	AreaSqm     float32  `json:"area_sqm"`
}

type UpdateRoomStatusRequest struct {
	Status string `json:"status"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type RoomResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Price       string    `json:"price"`
	RoomNumber  string    `json:"room_number"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	Id          string    `json:"id"`
	Description string    `json:"description,omitempty"`
	Title       string    `json:"title"`
	Amenities   []string  `json:"amenities"`
	Images      []string  `json:"images"`
	Capacity    int64     `json:"capacity"`
	Floor       int64     `json:"floor"`
	AreaSqm     float32   `json:"area_sqm"`
}

type RoomShortResponse struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	RoomNumber string   `json:"room_number"`
	Status     string   `json:"status"`
	Type       string   `json:"type"`
	Price      string   `json:"price"`
	Amenities  []string `json:"amenities"`
	Images     []string `json:"images"`
	Capacity   int64    `json:"capacity"`
	AreaSqm    float32  `json:"area_sqm"`
}

type RoomsResponse struct {
	Rooms []*RoomShortResponse `json:"rooms"`
}

// Booking DTOs
type CreateBookingRoomRequest struct {
	RoomId        string `json:"room_id"`
	PricePerNight string `json:"price_per_night"`
	Adults        uint32 `json:"adults"`
	Children      uint32 `json:"children"`
}

type CreateBookingRequest struct {
	CheckIn             time.Time                   `json:"check_in"`
	CheckOut            time.Time                   `json:"check_out"`
	HotelId             string                      `json:"hotel_id"`
	GuestName           string                      `json:"guest_name"`
	GuestEmail          string                      `json:"guest_email,omitempty"`
	GuestPhone          string                      `json:"guest_phone,omitempty"`
	Currency            string                      `json:"currency"`
	ExpectedTotalAmount string                      `json:"expected_total_amount"`
	Rooms               []*CreateBookingRoomRequest `json:"rooms"`
	UserId              int64                       `json:"user_id"`
}

type BookingRoomResponse struct {
	Id            string `json:"id"`
	RoomId        string `json:"room_id"`
	PricePerNight string `json:"price_per_night"`
	Adults        uint32 `json:"adults"`
	Children      uint32 `json:"children"`
}

type BookingResponse struct {
	CheckOut            time.Time              `json:"check_out"`
	UpdatedAt           time.Time              `json:"updated_at"`
	CreatedAt           time.Time              `json:"created_at"`
	CheckIn             time.Time              `json:"check_in"`
	GuestName           string                 `json:"guest_name"`
	Status              string                 `json:"status"`
	Id                  string                 `json:"id"`
	GuestEmail          string                 `json:"guest_email,omitempty"`
	GuestPhone          string                 `json:"guest_phone,omitempty"`
	Currency            string                 `json:"currency"`
	ExpectedTotalAmount string                 `json:"expected_total_amount"`
	FinalTotalAmount    string                 `json:"final_total_amount,omitempty"`
	HotelId             string                 `json:"hotel_id"`
	BookingRooms        []*BookingRoomResponse `json:"booking_rooms"`
	UserId              int64                  `json:"user_id"`
}

type BookingShortResponse struct {
	CheckIn             time.Time              `json:"check_in"`
	CheckOut            time.Time              `json:"check_out"`
	GuestName           string                 `json:"guest_name"`
	HotelId             string                 `json:"hotel_id"`
	Status              string                 `json:"status"`
	Id                  string                 `json:"id"`
	GuestEmail          string                 `json:"guest_email,omitempty"`
	GuestPhone          string                 `json:"guest_phone,omitempty"`
	Currency            string                 `json:"currency"`
	ExpectedTotalAmount string                 `json:"expected_total_amount"`
	FinalTotalAmount    string                 `json:"final_total_amount,omitempty"`
	BookingRooms        []*BookingRoomResponse `json:"booking_rooms"`
	UserId              int64                  `json:"user_id"`
}

type BookingsResponse struct {
	Bookings []*BookingShortResponse `json:"bookings"`
}

// Conversion functions
func TokenResponseFromProto(resp *userv1.RegisterUserResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}
	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromLoginProto(resp *userv1.LoginUserResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}
	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func TokenResponseFromRefreshProto(resp *userv1.RefreshTokenResponse) *TokenResponse {
	if resp == nil || resp.Tokens == nil {
		return nil
	}
	return &TokenResponse{
		Access:  resp.Tokens.Access,
		Refresh: resp.Tokens.Refresh,
	}
}

func UserResponseFromProto(user *userv1.User) *UserResponse {
	if user == nil {
		return nil
	}
	resp := &UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role.String(),
		IsActive: user.IsActive,
	}
	if user.Username != nil {
		resp.Username = *user.Username
	}
	if user.CreatedAt != nil {
		resp.CreatedAt = user.CreatedAt.AsTime()
	}
	if user.UpdatedAt != nil {
		resp.UpdatedAt = user.UpdatedAt.AsTime()
	}
	return resp
}

func userShortResponseFromProto(user *userv1.UserShort) *UserResponse {
	if user == nil {
		return nil
	}
	resp := &UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role.String(),
		IsActive: user.IsActive,
	}
	if user.Username != nil {
		resp.Username = *user.Username
	}
	return resp
}

func UsersResponseFromProto(resp *userv1.GetUsersResponse) *UsersResponse {
	if resp == nil {
		return nil
	}
	users := make([]*UserResponse, len(resp.Users))
	for i, u := range resp.Users {
		users[i] = userShortResponseFromProto(u)
	}
	return &UsersResponse{Users: users}
}

func UpdateUserResponseFromProto(user *userv1.UpdateUser) *UserResponse {
	if user == nil {
		return nil
	}
	resp := &UserResponse{
		Email: user.Email,
	}
	if user.Username != "" {
		resp.Username = user.Username
	}
	return resp
}

func locationDTOFromProto(loc *hotelv1.Location) *LocationDTO {
	if loc == nil {
		return nil
	}
	return &LocationDTO{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}
}

func HotelResponseFromProto(hotel *hotelv1.CreateHotel) *HotelResponse {
	if hotel == nil {
		return nil
	}
	resp := &HotelResponse{
		Id:          hotel.Id,
		Title:       hotel.Title,
		HotelSlug:   hotel.HotelSlug,
		OwnerId:     hotel.OwnerId,
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.CreatedAt != nil {
		resp.CreatedAt = hotel.CreatedAt.AsTime()
	}
	if hotel.UpdatedAt != nil {
		resp.UpdatedAt = hotel.UpdatedAt.AsTime()
	}
	return resp
}

func HotelDetailResponseFromProto(resp *hotelv1.GetHotelResponse) *HotelResponse {
	if resp == nil || resp.Hotel == nil {
		return nil
	}
	hotel := resp.Hotel
	result := &HotelResponse{
		Id:          hotel.Id,
		Title:       hotel.Title,
		OwnerId:     hotel.OwnerId,
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.Rating != nil {
		result.Rating = *hotel.Rating
	}
	if hotel.CreatedAt != nil {
		result.CreatedAt = hotel.CreatedAt.AsTime()
	}
	if hotel.UpdatedAt != nil {
		result.UpdatedAt = hotel.UpdatedAt.AsTime()
	}
	return result
}

func hotelShortResponseFromProto(hotel *hotelv1.HotelShort) *HotelShortResponse {
	if hotel == nil {
		return nil
	}
	resp := &HotelShortResponse{
		Id:        hotel.Id,
		Title:     hotel.Title,
		HotelSlug: hotel.HotelSlug,
		OwnerId:   hotel.OwnerId,
		Address:   hotel.Address,
		Location:  locationDTOFromProto(hotel.Location),
	}
	if hotel.Rating != nil {
		resp.Rating = *hotel.Rating
	}
	return resp
}

func HotelsShortResponseFromProto(resp *hotelv1.GetHotelsResponse) *HotelsResponse {
	if resp == nil {
		return nil
	}
	hotels := make([]*HotelShortResponse, len(resp.Hotels))
	for i, h := range resp.Hotels {
		hotels[i] = hotelShortResponseFromProto(h)
	}
	return &HotelsResponse{Hotels: hotels}
}

func UpdateHotelResponseFromProto(hotel *hotelv1.UpdateHotel) *HotelResponse {
	if hotel == nil {
		return nil
	}
	result := &HotelResponse{
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.Location != nil {
		result.Location = locationDTOFromProto(hotel.Location)
	}
	return result
}

func UpdateHotelTitleResponseFromProto(hotel *hotelv1.UpdateHotelTitle) *HotelResponse {
	if hotel == nil {
		return nil
	}
	return &HotelResponse{
		Title:     hotel.Title,
		HotelSlug: hotel.HotelSlug,
	}
}

func RoomResponseFromProto(room *hotelv1.Room) *RoomResponse {
	if room == nil {
		return nil
	}
	resp := &RoomResponse{
		Id:         room.Id,
		Title:      room.Title,
		RoomNumber: room.RoomNumber,
		Status:     room.Status.String(),
		Type:       room.Type.String(),
		Price:      room.Price,
		Capacity:   room.Capacity,
		AreaSqm:    room.AreaSqm,
		Floor:      room.Floor,
		Amenities:  room.Amenities,
		Images:     room.Images,
	}
	if room.Description != nil {
		resp.Description = *room.Description
	}
	if room.CreatedAt != nil {
		resp.CreatedAt = room.CreatedAt.AsTime()
	}
	if room.UpdatedAt != nil {
		resp.UpdatedAt = room.UpdatedAt.AsTime()
	}
	return resp
}

func roomShortResponseFromProto(room *hotelv1.RoomShort) *RoomShortResponse {
	if room == nil {
		return nil
	}
	return &RoomShortResponse{
		Id:         room.Id,
		Title:      room.Title,
		RoomNumber: room.RoomNumber,
		Status:     room.Status.String(),
		Type:       room.Type.String(),
		Price:      room.Price,
		Capacity:   room.Capacity,
		AreaSqm:    room.AreaSqm,
		Amenities:  room.Amenities,
		Images:     room.Images,
	}
}

func RoomsShortResponseFromProto(resp *hotelv1.GetRoomsResponse) *RoomsResponse {
	if resp == nil {
		return nil
	}
	rooms := make([]*RoomShortResponse, len(resp.Rooms))
	for i, r := range resp.Rooms {
		rooms[i] = roomShortResponseFromProto(r)
	}
	return &RoomsResponse{Rooms: rooms}
}

func UpdateRoomResponseFromProto(room *hotelv1.UpdateRoom) *RoomResponse {
	if room == nil {
		return nil
	}
	return &RoomResponse{
		Title:       room.Title,
		Description: room.Description,
		RoomNumber:  room.RoomNumber,
		Type:        room.Type.String(),
		Price:       room.Price,
		Capacity:    room.Capacity,
		AreaSqm:     room.AreaSqm,
		Floor:       room.Floor,
		Amenities:   room.Amenities,
		Images:      room.Images,
	}
}

func bookingRoomResponseFromProto(room *bookingv1.BookingRoom) *BookingRoomResponse {
	if room == nil {
		return nil
	}
	return &BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}

func bookingRoomWithLockResponseFromProto(room *bookingv1.BookingRoomWithLock) *BookingRoomResponse {
	if room == nil {
		return nil
	}
	return &BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}

func BookingResponseFromProto(booking *bookingv1.Booking) *BookingResponse {
	if booking == nil {
		return nil
	}
	resp := &BookingResponse{
		Id:                  booking.Id,
		UserId:              booking.UserId,
		HotelId:             booking.HotelId,
		Status:              booking.Status.String(),
		GuestName:           booking.GuestName,
		Currency:            booking.Currency,
		ExpectedTotalAmount: booking.ExpectedTotalAmount,
		FinalTotalAmount:    booking.FinalTotalAmount,
	}
	if booking.CheckIn != nil {
		resp.CheckIn = booking.CheckIn.AsTime()
	}
	if booking.CheckOut != nil {
		resp.CheckOut = booking.CheckOut.AsTime()
	}
	if booking.GuestEmail != nil {
		resp.GuestEmail = *booking.GuestEmail
	}
	if booking.GuestPhone != nil {
		resp.GuestPhone = *booking.GuestPhone
	}
	if booking.CreatedAt != nil {
		resp.CreatedAt = booking.CreatedAt.AsTime()
	}
	if booking.UpdatedAt != nil {
		resp.UpdatedAt = booking.UpdatedAt.AsTime()
	}
	rooms := make([]*BookingRoomResponse, len(booking.BookingRooms))
	for i, r := range booking.BookingRooms {
		rooms[i] = bookingRoomWithLockResponseFromProto(r)
	}
	resp.BookingRooms = rooms
	return resp
}

func bookingShortResponseFromProto(booking *bookingv1.BookingShort) *BookingShortResponse {
	if booking == nil {
		return nil
	}
	resp := &BookingShortResponse{
		Id:                  booking.Id,
		UserId:              booking.UserId,
		HotelId:             booking.HotelId,
		Status:              booking.Status.String(),
		GuestName:           booking.GuestName,
		Currency:            booking.Currency,
		ExpectedTotalAmount: booking.ExpectedTotalAmount,
		FinalTotalAmount:    booking.FinalTotalAmount,
	}
	if booking.CheckIn != nil {
		resp.CheckIn = booking.CheckIn.AsTime()
	}
	if booking.CheckOut != nil {
		resp.CheckOut = booking.CheckOut.AsTime()
	}
	if booking.GuestEmail != nil {
		resp.GuestEmail = *booking.GuestEmail
	}
	if booking.GuestPhone != nil {
		resp.GuestPhone = *booking.GuestPhone
	}
	rooms := make([]*BookingRoomResponse, len(booking.BookingRooms))
	for i, r := range booking.BookingRooms {
		rooms[i] = bookingRoomResponseFromProto(r)
	}
	resp.BookingRooms = rooms
	return resp
}

func BookingsShortResponseFromProto(resp *bookingv1.GetBookingsResponse) *BookingsResponse {
	if resp == nil {
		return nil
	}
	bookings := make([]*BookingShortResponse, len(resp.Bookings))
	for i, b := range resp.Bookings {
		bookings[i] = bookingShortResponseFromProto(b)
	}
	return &BookingsResponse{Bookings: bookings}
}
