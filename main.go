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
		}
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

	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)
	userController := controller.NewUserController(service)

	hotelRepo := repository.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	hotelController := controller.NewHotelController(hotelService)

	roomRepo := repository.NewRoomRepository(db)
	roomService := services.NewRoomService(roomRepo)
	roomController := controller.NewRoomController(roomService)

	updateroomRepo := repository.NewRoomRepository(db)
	updateroomService := services.NewRoomService(updateroomRepo)
	updateroomController := controller.NewRoomController(updateroomService)

	deletehotelRepo := repository.NewHotelRepository(db)
	deletehotelService := services.NewHotelService(deletehotelRepo)
	deletehotelController := controller.NewHotelController(deletehotelService)

	deleteroomRepo := repository.NewRoomRepository(db)
	deleteroomService := services.NewRoomService(deleteroomRepo)
	deleteroomController := controller.NewRoomController(deleteroomService)

	http.HandleFunc("/user/signup", userController.SignUp)
	http.HandleFunc("/user/signin", userController.SignIn)
	http.HandleFunc("/admin/addhotel", hotelController.AddHotel)
	http.HandleFunc("/admin/addroom", roomController.AddRoom)
	http.HandleFunc("/admin/updateroom", updateroomController.UpdateRoom)
	http.HandleFunc("/admin/deletehotel", deletehotelController.DeletHotel)
	http.HandleFunc("/admin/deleteroom", deleteroomController.DeleteRoom)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
