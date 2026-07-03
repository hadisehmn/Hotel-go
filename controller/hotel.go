package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
	"log"
	"net/http"
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
