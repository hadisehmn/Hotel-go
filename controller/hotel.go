package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
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
		fmt.Println(" add hotel error :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hotel Added ")

}

func (hd *HotelController) DeletHotel(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hotel deleted successfully")
}
