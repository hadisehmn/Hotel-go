package models

type Hotel struct {
	ID           int
	HotelName    string
	Star         int
	AveragePrice float64
}

type DeleteHotel struct {
	ID int `json:"id"`
}

type HotelList struct {
	Star         *int     `json:"star"`
	AveragePrice *float64 `json:"averageprice"`
}
