package apperror

import "errors"

var (
	ErrUserExists    = errors.New("user already exists")
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")

	ErrHotelExists   = errors.New("hotel already exists")
	ErrHotelNotFound = errors.New("hotel not found")

	ErrRoomExists       = errors.New("room already exists")
	ErrRoomNotFound     = errors.New("room not found")
	ErrInvalidRoomCount = errors.New("invalid room count")

	ErrNotEnoughRooms = errors.New("not enough rooms available")

	ErrInvalidData     = errors.New("invalid data")
	ErrInvalidDate     = errors.New("invalid date")
	ErrInvalidCapacity = errors.New("invalid capacity")

	ErrPriceNotFound = errors.New("room price not found")
)
