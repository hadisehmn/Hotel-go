package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-practice/HOTEL/apperror"
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
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, apperror.ErrUserNotFound
		}

		return models.User{}, fmt.Errorf("find user: %w", err)
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
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var roomID int

	err = tx.QueryRow(
		`INSERT INTO rooms(hotel_id, room_type, price, total_rooms, capacity)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		room.HotelID,
		room.RoomType,
		room.Price,
		room.TotalRooms,
		room.Capacity,
	).Scan(&roomID)

	if err != nil {
		return fmt.Errorf("create room: %w", err)
	}
	for _, p := range room.Prices {
		_, err := tx.Exec(
			`INSERT INTO pricing_rules (room_id, guest_type, price)
			 VALUES ($1, $2, $3)`,
			roomID,
			p.GuestType,
			p.Price,
		)
		if err != nil {
			return fmt.Errorf("create pricing rule: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
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
	result, err := r.DB.Exec(
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

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated room: %w", err)
	}

	if rows == 0 {
		return apperror.ErrRoomNotFound
	}
	return nil
}
func (r *HotelRepository) DeleteHotel(deletehotel models.DeleteHotel) error {
	result, err := r.DB.Exec(
		"DELETE FROM hotels WHERE id=$1",
		deletehotel.ID,
	)

	if err != nil {
		return fmt.Errorf("delete hotel: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted hotel: %w", err)
	}

	if rows == 0 {
		return apperror.ErrHotelNotFound
	}

	return nil
}

func (r *RoomRepository) DeleteRoom(deleteroom models.DeleteRoom) error {
	result, err := r.DB.Exec(
		"DELETE FROM rooms WHERE id=$1",
		deleteroom.ID,
	)

	if err != nil {
		return fmt.Errorf("delete room: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted room: %w", err)
	}

	if rows == 0 {
		return apperror.ErrRoomNotFound
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
		return nil, fmt.Errorf("query hotels: %w", err)
	}
	defer result.Close()

	for result.Next() {
		var h models.Hotel
		err := result.Scan(&h.ID, &h.HotelName, &h.Star, &h.AveragePrice)
		if err != nil {
			return nil, fmt.Errorf("scan hotel: %w", err)
		}
		hotels = append(hotels, h)

	}
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("iterate hotels: %w", err)
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
		return nil, fmt.Errorf("query rooms: %w", err)
	}
	defer result.Close()

	for result.Next() {
		var room models.Room

		err := result.Scan(
			&room.ID,
			&room.HotelID,
			&room.RoomType,
			&room.Price,
			&room.TotalRooms,
			&room.Capacity,
		)

		if err != nil {
			return nil, fmt.Errorf("scan room: %w", err)
		}

		priceRows, err := r.DB.Query(
			`SELECT id, room_id, guest_type, price 
			 FROM pricing_rules 
			 WHERE room_id = $1`,
			room.ID,
		)

		if err != nil {
			return nil, fmt.Errorf("get room prices: %w", err)
		}

		for priceRows.Next() {
			var p models.RoomPrice
			err := priceRows.Scan(
				&p.ID,
				&p.RoomID,
				&p.GuestType,
				&p.Price,
			)
			if err != nil {
				priceRows.Close()
				return nil, fmt.Errorf("scan room price: %w", err)
			}

			room.Prices = append(room.Prices, p)
		}
		priceRows.Close()
		rooms = append(rooms, room)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("iterate rooms: %w", err)
	}

	return rooms, nil
}

func (r *RoomRepository) FindRoomByID(RoomID int) (models.Room, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
			return models.Room{}, apperror.ErrRoomNotFound
		}

		return models.Room{}, fmt.Errorf("find room: %w", err)
	}
	return room, nil

}
func calculateTotalPrice(prices []models.RoomPrice, guests []models.GuestType) (float64, []models.GuestPriceDetail, error) {
	var total float64
	var guestPrices []models.GuestPriceDetail

	for _, guest := range guests {
		found := false
		for _, p := range prices {
			if guest == p.GuestType {

				total += p.Price
				guestPrices = append(guestPrices, models.GuestPriceDetail{
					GuestType: p.GuestType,
					Price:     p.Price,
				})
				found = true
				break
			}
		}
		if !found {
			return 0, nil, apperror.ErrPriceNotFound
		}
	}

	return total, guestPrices, nil
}

func (r *BookingRepository) BookRoom(UserID int, req models.BookRoomRequest, room models.Room) (models.Booking, []models.GuestPriceDetail, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var reservedRooms int

	err = tx.QueryRow(
		"SELECT COALESCE(SUM(room_count), 0)FROM bookings WHERE room_id = $1 AND check_out > $2 AND check_in < $3",
		req.RoomID,
		req.CheckIn,
		req.CheckOut,
	).Scan(&reservedRooms)
	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("get reserved rooms: %w", err)
	}
	availableRooms := room.TotalRooms - reservedRooms

	if availableRooms < req.RoomCount {
		return models.Booking{}, nil, apperror.ErrNotEnoughRooms
	}

	var prices []models.RoomPrice
	rows, err := tx.Query(
		`SELECT guest_type, price FROM pricing_rules WHERE room_id = $1`,
		req.RoomID,
	)
	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("get room prices: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var price models.RoomPrice

		if err := rows.Scan(&price.GuestType, &price.Price); err != nil {
			return models.Booking{}, nil, fmt.Errorf("scan room prices: %w", err)
		}
		prices = append(prices, price)
	}
	if err := rows.Err(); err != nil {
		return models.Booking{}, nil, fmt.Errorf("iterate room prices: %w", err)
	}

	oneNightPrice, guestPrices, err := calculateTotalPrice(prices, req.Guests)
	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("calculate total price: %w", err)
	}
	nights := int(req.CheckOut.Sub(req.CheckIn).Hours() / 24)
	TotalPrice := oneNightPrice * float64(nights) * float64(req.RoomCount)

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
		TotalPrice,
	).Scan(
		&booking.ID,
		&booking.CreatedAt,
	)

	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("insert booking: %w", err)
	}

	booking.RoomID = req.RoomID
	booking.RoomCount = req.RoomCount
	booking.CheckIn = req.CheckIn
	booking.CheckOut = req.CheckOut
	booking.GuestCount = len(req.Guests)
	booking.TotalPrice = TotalPrice
	// var guestPrices []models.GuestPriceDetail

	// for _, p := range prices {
	// 	guestPrices = append(guestPrices, models.GuestPriceDetail{
	// 		GuestType: p.GuestType,
	// 		Price:     p.Price,
	// 	})
	// }

	err = tx.Commit()
	if err != nil {
		return models.Booking{}, nil, fmt.Errorf("commit transaction: %w", err)
	}

	return booking, guestPrices, nil
}
