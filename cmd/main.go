package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"

	repository "github.com/SahilBheke25/ResourceSharingApplication/internal/repository"

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

	mux := http.DefaultServeMux

	// Routing
	mux.HandleFunc("POST /equipments", equipmentHandler.CreateEquipmentHandler)
	mux.HandleFunc("GET /equipments", equipmentHandler.ListEquipmentHandler)

	fmt.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
