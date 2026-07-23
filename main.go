package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"go-practice/HOTEL/controller"
	"go-practice/HOTEL/controller/middleware"
	"go-practice/HOTEL/repository"
	"go-practice/HOTEL/services"
)

func main() {

	//Method yml
	/*
		file, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		var cfg config.Config
		err = yaml.Unmarshal(file, &cfg)
		if err != nil {
			log.Fatal(err)

		db, err := sql.Open("postgres", cfg.Database.URI)
		if err != nil {
			log.Fatal(err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("DB connected:", cfg.Database.Name)
	*/

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseURI := os.Getenv("DATABASE_URI")

	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	fmt.Println("DB connected successfully")

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	hotelRepo := repository.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	hotelController := controller.NewHotelController(hotelService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := services.NewRoomService(roomRepo)
	roomController := controller.NewRoomController(roomService)

	bookingRepo := repository.NewBookingRepository(db)
	bookingService := services.NewBookingService(roomRepo, bookingRepo)
	bookingController := controller.NewBookingController(bookingService)

	http.HandleFunc("/user/signup", userController.SignUp)
	http.HandleFunc("/user/signin", userController.SignIn)
	http.HandleFunc("/user/hotellist", hotelController.HotelsList)
	http.HandleFunc("/user/roomlist", roomController.RoomList)

	http.Handle("/admin/addhotel",
		middleware.Authentication(
			middleware.AdminOnly(
				http.HandlerFunc(hotelController.AddHotel),
			),
		),
	)

	http.Handle("/admin/addroom",
		middleware.Authentication(
			middleware.AdminOnly(
				http.HandlerFunc(roomController.AddRoom),
			),
		),
	)

	http.Handle("/admin/updateroom",
		middleware.Authentication(
			middleware.AdminOnly(
				http.HandlerFunc(roomController.UpdateRoom),
			),
		),
	)

	http.Handle("/admin/deletehotel",
		middleware.Authentication(
			middleware.AdminOnly(
				http.HandlerFunc(hotelController.DeleteHotel),
			),
		),
	)

	http.Handle("/admin/deleteroom",
		middleware.Authentication(
			middleware.AdminOnly(
				http.HandlerFunc(roomController.DeleteRoom),
			),
		),
	)

	http.Handle(
		"/user/bookroom",
		middleware.Authentication(
			http.HandlerFunc(bookingController.BookRoom),
		),
	)

	http.Handle(
		"/user/bookinglist",
		middleware.Authentication(
			http.HandlerFunc(bookingController.GetBookingList),
		),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
