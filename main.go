package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"go-practice/HOTEL/controller"
	"go-practice/HOTEL/repository"
	"go-practice/HOTEL/services"
)

func main() {
	db, err := sql.Open("postgres",
		"postgres://postgres:123456@localhost:5436/hotel?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	fmt.Println("DB connected successfully")

	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)
	userController := controller.NewUserController(service)

	hotelRepo := repository.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	hotelController := controller.NewHotelController(hotelService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := services.NewRoomService(roomRepo)
	roomController := controller.NewRoomController(roomService)

	http.HandleFunc("/user/signup", userController.SignUp)
	http.HandleFunc("/user/signin", userController.SignIn)
	http.HandleFunc("/admin/addhotel", hotelController.AddHotel)
	http.HandleFunc("/admin/addroom", roomController.AddRoom)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
