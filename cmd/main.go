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

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	rentalService := rental.NewService(rentalRepo, equipmentService)
	rentalHandler := rental.NewHandler(rentalService)

	router := http.DefaultServeMux

	router.HandleFunc("POST /user/login", userHandler.Login)
	router.HandleFunc("POST /user/register", userHandler.Register)
	// router.HandleFunc("GET /user/{user_id}/edit-profile")

	router.HandleFunc("POST /equipments", equipmentHandler.CreateEquipmentHandler)
	router.HandleFunc("GET /equipments", equipmentHandler.ListEquipmentHandler)
	router.HandleFunc("DELETE /equipments/{equipment_id}", equipmentHandler.DeleteEquipmentHandler)
	router.HandleFunc("PUT /equipments/{equipment_id}", equipmentHandler.UpdateEquipmentHandler)
	router.HandleFunc("POST /users/{user_id}/equipments/{equip_id}/rent", rentalHandler.RentEquipment)
	router.HandleFunc("GET /users/{user_id}/equipments/lended", equipmentHandler.GetEquipmentsByUserIdHandler)
	router.HandleFunc("GET /equipments/{equipment_id}", equipmentHandler.EquipmentById)

	log.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", enableCORS(router)))

}
