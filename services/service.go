package services

import (
	"fmt"
	"go-practice/HOTEL/apperror"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

type HotelService struct {
	repo *repository.HotelRepository
}

type RoomService struct {
	repo *repository.RoomRepository
}
type BookingService struct {
	roomRepo    *repository.RoomRepository
	bookingRepo *repository.BookingRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func NewHotelService(repo *repository.HotelRepository) *HotelService {
	return &HotelService{
		repo: repo,
	}
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func NewBookingService(roomRepo *repository.RoomRepository, bookingRepo *repository.BookingRepository) *BookingService {
	return &BookingService{
		roomRepo:    roomRepo,
		bookingRepo: bookingRepo}
}

func (s *UserService) SignUp(u models.User) error {

	exists, err := s.repo.ExistsByName(u.Name)
	if err != nil {
		return fmt.Errorf("check user exists: %w", err)
	}

	if exists {
		return fmt.Errorf("signup: %w", apperror.ErrUserExists)
	}
	u.Role = models.UserRoleUser

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	u.Password = string(hashed)

	if err := s.repo.CreateUser(u); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (s *UserService) SignIn(u models.User) (models.User, error) {
	var emptyUser models.User

	user, err := s.repo.FindByName(u.Name)
	if err != nil {
		return emptyUser, fmt.Errorf("signin: %w", apperror.ErrUserNotFound)
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(u.Password))

	if err != nil {
		return emptyUser, fmt.Errorf("signin: %w", apperror.ErrWrongPassword)
	}
	return user, nil
}

func (s *HotelService) AddHotel(h models.Hotel) error {

	exist, err := s.repo.ExistsHotel(h.HotelName)
	if err != nil {
		return fmt.Errorf("check hotel exists: %w", err)
	}

	if exist {
		return fmt.Errorf("add hotel: %w", apperror.ErrHotelExists)
	}
	if err := s.repo.CreateHotel(h); err != nil {
		return fmt.Errorf("create hotel: %w", err)
	}

	return nil
}

func (s *RoomService) AddRoom(room models.Room) error {

	if room.Capacity <= 0 {
		return fmt.Errorf("add room: %w", apperror.ErrInvalidCapacity)
	}

	if room.TotalRooms <= 0 {
		return fmt.Errorf("add room: %w", apperror.ErrInvalidRoomCount)
	}

	exist, err := s.repo.ExistRoom(room.HotelID, room.RoomType)
	if err != nil {
		return fmt.Errorf("check room exists: %w", err)
	}
	if exist {
		return fmt.Errorf("add room: %w", apperror.ErrRoomExists)
	}

	if err := s.repo.CreateRoom(room); err != nil {
		return fmt.Errorf("create room: %w", err)
	}

	return nil
}

func (s *RoomService) UpdateRoom(id int, roomup models.UpdateRoom) error {
	// fmt.Printf("roomup: %+v\n", roomup)
	if err := s.repo.UpdateRoom(id, roomup); err != nil {
		return fmt.Errorf("update room: %w", err)
	}

	return nil
}

func (s *HotelService) DeleteHotel(deletehotel models.DeleteHotel) error {
	if err := s.repo.DeleteHotel(deletehotel); err != nil {
		return fmt.Errorf("delete hotel: %w", err)
	}

	return nil
}

func (s *RoomService) DeleteRoom(deleteroom models.DeleteRoom) error {
	if err := s.repo.DeleteRoom(deleteroom); err != nil {
		return fmt.Errorf("delete room: %w", err)
	}

	return nil
}

func (s *HotelService) HotelsList(filter models.HotelList) ([]models.Hotel, error) {
	hotels, err := s.repo.HotelsList(filter)
	if err != nil {
		return nil, fmt.Errorf("list hotels: %w", err)
	}
	return hotels, nil
}

func (s *RoomService) RoomList(filter models.RoomList) ([]models.Room, error) {
	rooms, err := s.repo.RoomList(filter)
	if err != nil {
		return nil, fmt.Errorf("list rooms: %w", err)
	}
	return rooms, nil

}

func (s *BookingService) BookRoom(UserID int, req models.BookRoomRequest) (models.Booking, error) {

	room, err := s.roomRepo.FindRoomByID(req.RoomID)
	if err != nil {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrRoomNotFound)
	}
	if req.RoomCount <= 0 {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidData)
	}
	if req.CheckOut.Before(req.CheckIn) || req.CheckOut.Equal(req.CheckIn) {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidDate)
	}
	if req.CheckIn.Before(time.Now()) {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidDate)
	}
	if len(req.Guests) == 0 {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidData)
	}
	if len(req.Guests) > room.Capacity*req.RoomCount {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidCapacity)
	}
	if room.Capacity < len(req.Guests) {
		return models.Booking{}, fmt.Errorf("book room: %w", apperror.ErrInvalidCapacity)
	}
	booking, err := s.bookingRepo.BookRoom(UserID, req, room)
	if err != nil {
		return models.Booking{}, fmt.Errorf("create booking: %w", err)
	}
	return booking, nil
}
