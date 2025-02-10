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

	equipmentRepo := repository.NewEquipmentStore(db)
	equipmentService := equipment.NewService(equipmentRepo)
	equipmentHandler := equipment.NewHandler(equipmentService)

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /equipments", equipmentHandler.CreateEquipmentHandler)
	mux.HandleFunc("GET /equipments", equipmentHandler.ListEquipmentHandler)
	mux.HandleFunc("PUT /equipments/{equipment_id}", equipmentHandler.UpdateEquipmentHandler)

	fmt.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
