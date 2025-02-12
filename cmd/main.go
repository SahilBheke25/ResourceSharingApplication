package main

import (
	"log"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/rental"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/user"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	// Creating DB connection
	db := repository.InitializeDatabase()
	defer db.Close()

	// Initialize Dependencies
	equipmentRepo := repository.NewEquipmentStore(db)
	equipmentService := equipment.NewService(equipmentRepo)
	equipmentHandler := equipment.NewHandler(equipmentService)

	userRepo := repository.NewUserStorer(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	rentalRepo := repository.NewRentalStore(db)
	rentalService := rental.NewService(rentalRepo)
	rentalHandler := rental.NewHandler(rentalService)

	router := http.DefaultServeMux

	router.HandleFunc("POST /login", userHandler.VerifyUserHandler)
	router.HandleFunc("POST /register", userHandler.RegisterUserHandler)
	router.HandleFunc("POST /equipments", equipmentHandler.CreateEquipmentHandler)
	router.HandleFunc("GET /equipments", equipmentHandler.ListEquipmentHandler)
	router.HandleFunc("DELETE /equipments/{equipment_id}", equipmentHandler.DeleteEquipmentHandler)
	router.HandleFunc("PUT /equipments/{equipment_id}", equipmentHandler.UpdateEquipmentHandler)
	router.HandleFunc("POST /users/{user_id}/equipments/{equip_id}/rent", rentalHandler.RentEquipment)
	router.HandleFunc("GET /equipments/{user_id}", equipmentHandler.GetEquipmentsByUserIdHandler)

	log.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
