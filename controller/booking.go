package controller

import (
	"encoding/json"
	"errors"
	"go-practice/HOTEL/apperror"
	"go-practice/HOTEL/models"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "user id missing", http.StatusUnauthorized)
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

	booking, guestPrices, err := br.service.BookRoom(userID, req)
	if err != nil {
		log.Printf("Booking room failed: %v", err)

		switch {

		case errors.Is(err, apperror.ErrRoomNotFound):
			http.Error(w, "Room not found", http.StatusNotFound)

		case errors.Is(err, apperror.ErrNotEnoughRooms):
			http.Error(w, "Not enough rooms available", http.StatusConflict)

		case errors.Is(err, apperror.ErrInvalidCapacity):
			http.Error(w, "Room capacity exceeded", http.StatusBadRequest)

		case errors.Is(err, apperror.ErrInvalidData):
			http.Error(w, "Invalid booking information", http.StatusBadRequest)

		case errors.Is(err, apperror.ErrInvalidDate):
			http.Error(w, "Invalid check-in/check-out date", http.StatusBadRequest)

		case errors.Is(err, apperror.ErrPriceNotFound):
			http.Error(w, "Room price not found", http.StatusNotFound)

		default:
			log.Printf("BookRoom: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := models.BookingResponse{
		Message:     "Room booked successfully",
		Booking:     booking,
		GuestPrices: guestPrices,
	}
	json.NewEncoder(w).Encode(response)
}

func (br *BookingController) GetBookingList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "user id missing", http.StatusUnauthorized)
		return
	}
	bookings, err := br.service.GetBookingList(userID)
	if err != nil {
		log.Printf("GetBookingList: %v", err)

		switch {
		case errors.Is(err, apperror.ErrBookingNotFound):
			http.Error(w, "You don't have any bookings yet", http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		log.Printf("encode bookings response: %v", err)
	}
}
