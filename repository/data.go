package repository

import (
	"database/sql"
	"fmt"
	"go-practice/HOTEL/models"
)

type UserRepository struct {
	DB *sql.DB
}
type HotelRepository struct {
	DB *sql.DB
}

type RoomRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func NewHotelRepository(db *sql.DB) *HotelRepository {
	return &HotelRepository{
		DB: db,
	}
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}

}

func (r *UserRepository) CreateUser(u models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(name, phone, password_hash) VALUES ($1, $2, $3)",
		u.Name,
		u.Phone,
		u.Password,
	)
	return err
}

func (r *UserRepository) ExistsByName(name string) (bool, error) {
	var exists bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE name=$1)",
		name,
	).Scan(&exists)

	return exists, err
}

func (r *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT name, password_hash FROM users WHERE name = $1", name,
	).Scan(
		&user.Name,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil

}

func (r *HotelRepository) CreateHotel(h models.Hotel) error {

	_, err := r.DB.Exec(
		"INSERT INTO hotels(hotel_name, star, average_price) VALUES ($1, $2, $3)",
		h.HotelName,
		h.Star,
		h.AveragePrice,
	)
	return err

}

func (r *HotelRepository) ExistsHotel(HotelName string) (bool, error) {
	var Exist bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM hotels WHERE hotel_name=$1)",
		HotelName,
	).Scan(&Exist)
	return Exist, err

}

func (r *RoomRepository) CreateRoom(room models.Room) error {

	_, err := r.DB.Exec(
		"INSERT INTO rooms(hotel_id, room_name, room_type, price)VALUES ($1, $2, $3, $4)",
		room.HotelID,
		room.RoomName,
		room.RoomType,
		room.Price,
	)
	return err

}

func (r *RoomRepository) ExistRoom(HotelID int, RoomName string) (bool, error) {
	var Exist bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM rooms WHERE hotel_id=$1 AND room_name=$2)",
		HotelID,
		RoomName,
	).Scan(&Exist)
	return Exist, err
}

func (r *RoomRepository) UpdateRoom(id int, roomup models.UpdateRoom) error {
	_, err := r.DB.Exec(
		"UPDATE rooms SET room_name = $1, room_type = $2, price = $3 WHERE id = $4",
		roomup.RoomName,
		roomup.RoomType,
		roomup.Price,
		id,
	)
	return err

}

func (r *HotelRepository) DeleteHotel(id int, hotelName string) error {
	_, err := r.DB.Exec(
		"DELETE FROM hotels WHERE id=$1 AND hotel_name=$2",
		id,
		hotelName,
	)
	if err != nil {
		return err
	}
	return nil
}

// func (r *RoomRepository) FindById(id int) (models.UpdateRoom, error) {

// 	var update models.UpdateRoom
// 	err := r.DB.QueryRow(
// 		"SELECT id FROM rooms WHERE id = $1",
// 		id,
// 	).Scan(
// 		&update.ID,
// 	)
// 	if err == sql.ErrNoRows {
// 		return update, fmt.Errorf("room not found")
// 	}

// 	return update, err

// }
