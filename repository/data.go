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
type BookingRepository struct {
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
func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		DB: db,
	}

}

func (r *UserRepository) CreateUser(u models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(name, phone, password_hash , role) VALUES ($1, $2, $3 ,$4 )",
		u.Name,
		u.Phone,
		u.Password,
		u.Role,
	)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *UserRepository) ExistsByName(name string) (bool, error) {
	var exists bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE name=$1)",
		name,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("check user exists: %w", err)
	}

	return exists, nil
}

func (r *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT id, name, password_hash, role FROM users WHERE name = $1", name,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("find user by name: %w", err)
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
	if err != nil {
		return fmt.Errorf("create hotel: %w", err)
	}
	return nil
}

func (r *HotelRepository) ExistsHotel(HotelName string) (bool, error) {
	var Exist bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM hotels WHERE hotel_name=$1)",
		HotelName,
	).Scan(&Exist)
	if err != nil {
		return false, fmt.Errorf("check hotel exists: %w", err)
	}

	return Exist, nil

}

func (r *RoomRepository) CreateRoom(room models.Room) error {

	_, err := r.DB.Exec(
		"INSERT INTO rooms(hotel_id, room_type, price, total_rooms, capacity) VALUES ($1, $2, $3, $4, $5)",
		room.HotelID,
		room.RoomType,
		room.Price,
		room.TotalRooms,
		room.Capacity,
	)
	if err != nil {
		return fmt.Errorf("create room: %w", err)
	}
	return nil
}

func (r *RoomRepository) ExistRoom(HotelID int, roomType models.RoomType) (bool, error) {
	var Exist bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM rooms WHERE hotel_id=$1 AND room_type=$2)",
		HotelID,
		roomType,
	).Scan(&Exist)
	if err != nil {
		return false, fmt.Errorf("check room exists: %w", err)
	}
	return Exist, nil
}

func (r *RoomRepository) UpdateRoom(id int, roomup models.UpdateRoom) error {
	_, err := r.DB.Exec(
		"UPDATE rooms SET room_type=$1, price=$2, total_rooms=$3 , capacity=$4 WHERE id = $5",
		roomup.RoomType,
		roomup.Price,
		roomup.TotalRooms,
		roomup.Capacity,
		id,
	)
	if err != nil {
		return fmt.Errorf("update room: %w", err)
	}
	return nil
}
func (r *HotelRepository) DeleteHotel(deletehotel models.DeleteHotel) error {
	_, err := r.DB.Exec(
		"DELETE FROM hotels WHERE id=$1",
		deletehotel.ID,
	)
	if err != nil {
		return fmt.Errorf("delete hotel: %w", err)
	}
	return nil
}

func (r *RoomRepository) DeleteRoom(deleteroom models.DeleteRoom) error {
	_, err := r.DB.Exec(
		"DELETE FROM rooms WHERE id=$1",
		deleteroom.ID,
	)
	if err != nil {
		return fmt.Errorf("delete room: %w", err)
	}
	return nil
}

func (r *HotelRepository) HotelsList(filter models.HotelList) ([]models.Hotel, error) {
	var hotels []models.Hotel

	query := "SELECT id, hotel_name, star, average_price FROM hotels WHERE 1=1"
	params := []any{}
	i := 1

	if filter.Star != nil {
		query += fmt.Sprintf(" AND star = $%d", i)
		params = append(params, *filter.Star)
		i++

	}
	if filter.AveragePrice != nil {
		query += fmt.Sprintf(" AND average_price >= $%d", i)
		params = append(params, *filter.AveragePrice)
		i++
	}
	result, err := r.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var h models.Hotel
		err := result.Scan(&h.ID, &h.HotelName, &h.Star, &h.AveragePrice)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, h)

	}
	return hotels, nil

}

func (r *RoomRepository) RoomList(filter models.RoomList) ([]models.Room, error) {
	var rooms []models.Room
	query := "SELECT id, hotel_id, room_type, price, total_rooms, capacity FROM rooms WHERE 1=1"
	params := []any{}
	i := 1

	if filter.Price != nil {
		query += fmt.Sprintf(" AND price >= $%d", i)
		params = append(params, *filter.Price)
		i++
	}
	if filter.RoomType != nil {
		query += fmt.Sprintf(" AND room_type ILIKE $%d", i)
		params = append(params, *filter.RoomType)
		i++

	}
	result, err := r.DB.Query(query, params...)
	if err != nil {
		return nil, err

	}

	for result.Next() {
		var r models.Room
		err := result.Scan(&r.ID, &r.HotelID, &r.RoomType, &r.Price, &r.TotalRooms, &r.Capacity)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, r)

	}
	return rooms, nil

}

func (r *RoomRepository) FindRoomById(RoomID int) (models.Room, error) {
	var room models.Room

	err := r.DB.QueryRow(
		`SELECT id, hotel_id, room_type, price, total_rooms, capacity FROM rooms WHERE id = $1`,
		RoomID,
	).Scan(
		&room.ID,
		&room.HotelID,
		&room.RoomType,
		&room.Price,
		&room.TotalRooms,
		&room.Capacity,
	)

	if err != nil {
		return models.Room{}, fmt.Errorf(" %w", err)
	}
	return room, nil

}

func (r *BookingRepository) BookRoom(UserID int, req models.BookRoomRequest, room models.Room) (models.Booking, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return models.Booking{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var booking models.Booking
	err = tx.QueryRow(
		`INSERT INTO bookings (
			user_id,
			room_id,
			room_count,
			check_in,
			check_out,
			guest_count,
			total_price
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id, created_at`,
		UserID,
		req.RoomID,
		req.RoomCount,
		req.CheckIn,
		req.CheckOut,
		len(req.Guests),
		room.Price*float64(req.RoomCount),
	).Scan(
		&booking.ID,
		&booking.CreatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("insert booking: %w", err)
	}
	_, err = tx.Exec(
		`UPDATE rooms
		SET total_rooms = total_rooms - $1
		WHERE id = $2`,
		req.RoomCount,
		req.RoomID,
	)
	if err != nil {
		return models.Booking{}, fmt.Errorf("update room: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Booking{}, fmt.Errorf("commit transaction: %w", err)
	}

	return booking, nil
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
