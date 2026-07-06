package services

import (
	"fmt"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/repository"

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

func (s *UserService) SignUp(u models.User) error {

	exists, err := s.repo.ExistsByName(u.Name)
	if err != nil {
		return fmt.Errorf("signup: %w", err)
	}

	if exists {
		return fmt.Errorf("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	u.Password = string(hashed)

	if err := s.repo.CreateUser(u); err != nil {
		return fmt.Errorf("signup: %w", err)
	}

	return nil
}

func (s *UserService) SignIn(u models.User) (models.User, error) {
	var emptyUser models.User

	user, err := s.repo.FindByName(u.Name)
	if err != nil {
		return emptyUser, fmt.Errorf("user not found")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(u.Password))

	if err != nil {
		return emptyUser, fmt.Errorf("wrong password")
	}
	return user, nil
}

func (s *HotelService) AddHotel(h models.Hotel) error {

	exist, err := s.repo.ExistsHotel(h.HotelName)
	if err != nil {
		return fmt.Errorf("add hotel: %w", err)
	}

	if exist {
		return fmt.Errorf("hotel already exists")
	}

	if err := s.repo.CreateHotel(h); err != nil {
		return fmt.Errorf("add hotel: %w", err)
	}

	return nil
}

func (s *RoomService) AddRoom(room models.Room) error {
	exist, err := s.repo.ExistRoom(room.HotelID, room.RoomName)
	if err != nil {
		return fmt.Errorf("add room: %w", err)
	}

	if exist {
		return fmt.Errorf("Room already exists")
	}

	if err := s.repo.CreateRoom(room); err != nil {
		return fmt.Errorf("add room: %w", err)
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
