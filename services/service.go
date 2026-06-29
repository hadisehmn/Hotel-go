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

func (s *UserService) SignUp(u models.User) error {

	exists, err := s.repo.ExistsByName(u.Name)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)

	return s.repo.CreateUser(u)
}

func (s *UserService) SignIn(u models.User) error {

	user, err := s.repo.FindByName(u.Name)
	if err != nil {
		return err

	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(u.Password))

	if err != nil {
		return fmt.Errorf("wrong password")
	} else {
		fmt.Println("Login Successful")
		return nil
	}
}

func (s *HotelService) AddHotel(h models.Hotel) error {

	exist, err := s.repo.ExistsHotel(h.HotelName)
	if err != nil {
		return err
	}

	if exist {
		fmt.Println("Hotel existed")
		return fmt.Errorf("hotel already exists")
	}

	fmt.Println("Hotel added")
	return s.repo.CreateHotel(h)
}
