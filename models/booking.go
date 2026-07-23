package models

import (
	"time"
)

type Booking struct {
	ID         int       `json:"id"`
	UserID     int       `json:"-"`
	RoomID     int       `json:"room_id"`
	RoomCount  int       `json:"room_count"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	GuestCount int       `json:"guest_count"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type GuestPriceDetail struct {
	GuestType GuestType `json:"guest_type"`
	Price     float64   `json:"price"`
}
type BookingResponse struct {
	Message     string             `json:"message"`
	Booking     Booking            `json:"booking"`
	GuestPrices []GuestPriceDetail `json:"guest_prices"`
}

type BookingGuestPrice struct {
	ID        int
	BookingID int
	GuestType GuestType
	Price     float64
}

type BookingGuestPriceResponse struct {
	GuestType GuestType `json:"guest_type"`
	Price     float64   `json:"price"`
}

type BookingList struct {
	Booking     Booking                     `json:"booking"`
	HotelName   string                      `json:"hotel_name"`
	RoomType    RoomType                    `json:"room_type"`
	GuestPrices []BookingGuestPriceResponse `json:"guest_prices"`
}
