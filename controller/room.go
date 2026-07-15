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

type RoomController struct {
	service *services.RoomService
}

func NewRoomController(service *services.RoomService) *RoomController {
	return &RoomController{
		service: service,
	}

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
		log.Printf("AddRoom failed: %v", err)

		switch err.Error() {

		case "room already exists":
			http.Error(w, "Room already exists", http.StatusConflict)
			return

		default:
			http.Error(w, "Room already exists", http.StatusInternalServerError)

		}
		return
	}
	fmt.Fprintln(w, "Room added successfully")

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
		log.Printf("UpdateRoom failed: %v", err)

		switch err.Error() {
		case "room not found":
			http.Error(w, "Room not found", http.StatusNotFound)
			return
		default:
			http.Error(w, "Failed to update room", http.StatusInternalServerError)

		}
		return
	}
	fmt.Fprintln(w, "room updated ")

}

func (rd *RoomController) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var deleteroom models.DeleteRoom
	err := json.NewDecoder(r.Body).Decode(&deleteroom)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if deleteroom.ID == 0 {
		http.Error(w, "Missing room  id ", http.StatusBadRequest)
		return
	}
	err = rd.service.DeleteRoom(deleteroom)
	if err != nil {
		log.Printf("DeleteRoom failed: %v", err)
		http.Error(w, "Failed to delete room", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Room deleted successfully")
}

func (rl *RoomController) RoomList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var filter models.RoomList

	RoomTypeParam := r.URL.Query().Get("roomtype")
	priceParam := r.URL.Query().Get("price")

	if RoomTypeParam != "" {
		roomType := models.RoomType(RoomTypeParam)
		filter.RoomType = &roomType
	}

	if priceParam != "" {
		p, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			http.Error(w, "invalid price", http.StatusBadRequest)
			return
		}

		filter.Price = &p
	}

	list, err := rl.service.RoomList(filter)
	if err != nil {
		http.Error(w, "Failed to get rooms", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if len(list) == 0 {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "No rooms found",
		})
		return
	}

	json.NewEncoder(w).Encode(list)

}

// func (br *RoomController) BookRoom(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	roomIdParam := r.URL.Query().Get("room_id")

// 	if roomIdParam == "" {
// 		http.Error(w, "room_id is Require ", http.StatusBadRequest)
// 		return

// 	}
// 	roomID, err := strconv.Atoi(roomIdParam)

// 	if err != nil {
// 		http.Error(w, "invalid room_id", http.StatusBadRequest)
// 		return
// 	}

// 	var reserve models.BookRoomRequest

// 	err = json.NewDecoder(r.Body).Decode(&reserve)
// 	if err != nil {
// 		http.Error(w, "invalid body", http.StatusBadRequest)
// 		return
// 	}

// 	if reserve.CheckIn.IsZero() || reserve.CheckOut.IsZero() {
// 		http.Error(w, "check_in and check_out are required", http.StatusBadRequest)
// 		return

// 	}

// 	booking, err := br.service.BookRoom(roomID, reserve)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(booking)

// }
