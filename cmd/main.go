package main

import (
	"log"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"
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

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /login", userHandler.VerifyUserHandler)
	mux.HandleFunc("POST /register", userHandler.RegisterUserHandler)
	mux.HandleFunc("POST /equipments", equipmentHandler.CreateEquipmentHandler)
	mux.HandleFunc("GET /equipments", equipmentHandler.ListEquipmentHandler)
	mux.HandleFunc("DELETE /equipments/{equipment_id}", equipmentHandler.DeleteEquipmentHandler)
	mux.HandleFunc("PUT /equipments/{equipment_id}", equipmentHandler.UpdateEquipmentHandler)

	log.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
