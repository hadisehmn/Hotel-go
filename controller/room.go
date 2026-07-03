package controller

import (
	"encoding/json"
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/services"
	"log"
	"net/http"
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
