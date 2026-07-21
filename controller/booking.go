package controller

import (
	"encoding/json"
	"errors"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/repository"
	"go-practice/HOTEL/services"
	"log"
	"net/http"
)

type BookingController struct {
	service *services.BookingService
}

func NewBookingController(service *services.BookingService) *BookingController {
	return &BookingController{
		service: service,
	}

}

func (br *BookingController) BookRoom(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "user id missing", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.BookRoomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if req.RoomID == 0 {
		http.Error(w, "room_id is required", http.StatusBadRequest)
		return

	}

	if req.CheckIn.IsZero() || req.CheckOut.IsZero() {
		http.Error(w, "check_in and check_out time is required", http.StatusBadRequest)
		return

	}

	booking, err := br.service.BookRoom(userID, req)
	if err != nil {
		log.Printf("Booking room failed: %v", err)

		switch {
		case errors.Is(err, services.ErrRoomNotFound):
			http.Error(w, "Room not found", http.StatusNotFound)
			return

		case errors.Is(err, repository.ErrNotEnoughRooms):
			http.Error(w, "Not enough rooms available", http.StatusConflict)
			return

		case errors.Is(err, services.ErrInvalidCapacity):
			http.Error(w, "Room capacity exceeded", http.StatusBadRequest)
			return

		case errors.Is(err, services.ErrInvalidData):
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return

		case errors.Is(err, services.ErrInvalidDate):
			http.Error(w, "Invalid date", http.StatusBadRequest)
			return

		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}

	response := models.BookingResponse{
		Message: "Room booked successfully",
		Booking: booking,
	}

	json.NewEncoder(w).Encode(response)
}
