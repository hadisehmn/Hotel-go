package models

import "time"

type RoomType string

const (
	Standard RoomType = "standard"
	Single   RoomType = "single"
	Double   RoomType = "double"
)

type GuestType string

const (
	Adult  GuestType = "adult"
	Child  GuestType = "child"
	Infant GuestType = "infant"
)

type Room struct {
	ID       int      `json:"id"`
	HotelID  int      `json:"hotel_id"`
	RoomName string   `json:"room_name"`
	RoomType RoomType `json:"room_type"`
	Price    float64  `json:"price"`
	Capacity int      `json:"capacity"`
}

type UpdateRoom struct {
	ID       int      `json:"id"`
	RoomName string   `json:"room_name"`
	RoomType RoomType `json:"room_type"`
	Price    float64  `json:"price"`
	Capacity int      `json:"capacity"`
}
type DeleteRoom struct {
	ID int `json:"id"`
}

type RoomList struct {
	Price    *float64  `json:"price"`
	RoomType *RoomType `json:"room_type"`
}

type BookRoomRequest struct {
	RoomID   int         `json:"room_id"`
	UserID   int         `json:"user_id"`
	CheckIn  time.Time   `json:"check_in"`
	CheckOut time.Time   `json:"check_out"`
	Guests   []GuestType `json:"guests"`
}

// admin
