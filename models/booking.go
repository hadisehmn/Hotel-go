package models

import "time"

type Booking struct {
	ID         int
	UserID     int
	RoomID     int
	CheckIn    time.Time
	CheckOut   time.Time
	GuestCount int
	TotalPrice float64
	CreatedAt  time.Time
}
