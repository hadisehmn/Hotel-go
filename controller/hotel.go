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

type RoomController struct {
	service *services.RoomService
}

func NewHotelController(service *services.HotelService) *HotelController {
	return &HotelController{
		service: service,
	}
}

func NewRoomController(service *services.RoomService) *RoomController {
	return &RoomController{
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

func (rc *RoomController) AddRoom(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = rc.service.AddRoom(room)
	if err != nil {
		fmt.Println(" add room error :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "rooms Added ")

}

func (ru *RoomController) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var roomup models.UpdateRoom

	err := json.NewDecoder(r.Body).Decode(&roomup)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	err = ru.service.UpdateRoom(roomup.ID, roomup)
	if err != nil {
		fmt.Println(" update room error :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "room updated ")

}
