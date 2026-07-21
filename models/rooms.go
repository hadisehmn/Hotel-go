package models

import "time"

type RoomType string

const (
	Single RoomType = "single"
	Double RoomType = "double"
	Suite  RoomType = "suite"
)

type GuestType string

const (
	Adult  GuestType = "adult"
	Child  GuestType = "child"
	Infant GuestType = "infant"
)

type Room struct {
	ID         int      `json:"id"`
	HotelID    int      `json:"hotel_id"`
	RoomType   RoomType `json:"room_type"`
	Price      float64  `json:"price"`
	TotalRooms int      `json:"total_rooms"`
	Capacity   int      `json:"capacity"`
}

type UpdateRoom struct {
	ID         int      `json:"id"`
	RoomType   RoomType `json:"room_type"`
	Price      float64  `json:"price"`
	TotalRooms int      `json:"total_rooms"`
	Capacity   int      `json:"capacity"`
}
type DeleteRoom struct {
	ID int `json:"id"`
}

type RoomList struct {
	Price    *float64  `json:"price"`
	RoomType *RoomType `json:"room_type"`
}

type BookRoomRequest struct {
	RoomID    int         `json:"room_id"`
	RoomCount int         `json:"room_count"`
	CheckIn   time.Time   `json:"check_in"`
	CheckOut  time.Time   `json:"check_out"`
	Guests    []GuestType `json:"guests"`
}

type RoomPrice struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	GuestType GuestType `json:"guest_type"`
	Price     float64   `json:"price"`
}
