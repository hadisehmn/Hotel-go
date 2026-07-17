package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
	"log"
	"net/http"
	"strconv"
)

type HotelController struct {
	service *services.HotelService
}

func NewHotelController(service *services.HotelService) *HotelController {
	return &HotelController{
		service: service,
	}
}

func (hc *HotelController) AddHotel(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var h models.Hotel
	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = hc.service.AddHotel(h)
	if err != nil {
		log.Printf("AddHotel failed: %v", err)
		log.Printf("AddHotel failed: %v", err)

		if err.Error() == "hotel already exists" {
			http.Error(w, "Hotel already exists", http.StatusConflict)
			return
		}

		http.Error(w, "Failed to add hotel", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hotel Added ")

}

func (hd *HotelController) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var deletehotel models.DeleteHotel
	err := json.NewDecoder(r.Body).Decode(&deletehotel)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if deletehotel.ID == 0 {
		http.Error(w, "Missing hotel id ", http.StatusBadRequest)
		return
	}
	err = hd.service.DeleteHotel(deletehotel)
	if err != nil {
		log.Printf("DeleteHotel failed: %v", err)

		if err.Error() == "hotel not found" {
			http.Error(w, "Hotel not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete hotel", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hotel deleted successfully")
}

func (hl *HotelController) HotelsList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	var filter models.HotelList

	starParam := r.URL.Query().Get("star")
	priceParam := r.URL.Query().Get("averageprice")

	if starParam != "" {
		star, err := strconv.Atoi(starParam)
		if err != nil {
			http.Error(w, "invalid star", http.StatusBadRequest)
			return

		}
		filter.Star = &star
	}

	if priceParam != "" {
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			http.Error(w, "invalid average price", http.StatusBadRequest)
			return
		}
		filter.AveragePrice = &price
	}

	list, err := hl.service.HotelsList(filter)
	if err != nil {
		http.Error(w, "Failed to get hotels", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if len(list) == 0 {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "No hotels found",
		})
		return
	}

	json.NewEncoder(w).Encode(list)

}

// admin
