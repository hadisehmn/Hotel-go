package models

type RoomType string

const (
	Standard RoomType = "standard"
	Single   RoomType = "single"
	Double   RoomType = "double"
)

type Room struct {
	HotelID  int      `json:"hotel_id"`
	RoomName string   `json:"room_name"`
	RoomType RoomType `json:"room_type"`
	Price    float64  `json:"price"`
}

type UpdateRoom struct {
	ID       int      `json:"id"`
	RoomName string   `json:"room_name"`
	RoomType RoomType `json:"room_type"`
	Price    float64  `json:"price"`
}
type DeleteRoom struct {
	ID int `json:"id"`
}
