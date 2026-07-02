package models

type Hotel struct {
	ID           int
	HotelName    string
	Star         int
	AveragePrice int
}

type DeleteHotel struct {
	ID        int    `json:"id"`
	HotelName string `json:"hotel_name"`
}
