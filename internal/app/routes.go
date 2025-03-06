package app

import "net/http"

func InitializeRoutes(deps *Dependencies) http.Handler {

	router := http.DefaultServeMux

	// User Routes
	router.HandleFunc("POST /user/login", deps.userHandler.Login)
	router.HandleFunc("POST /user/register", deps.userHandler.Register)
	router.HandleFunc("GET /owner/equipment/{equipment_id}", deps.userHandler.EquipmentOwner)
	router.HandleFunc("GET /user/{user_id}", deps.userHandler.UserById)
	router.HandleFunc("PUT /user/edit-profile/{user_id}", deps.userHandler.UpdateProfile)

	// Equipment Routes
	router.HandleFunc("POST /equipments", deps.equipmentHandler.CreateEquipmentHandler)
	router.HandleFunc("POST /users/{user_id}/equipments/{equip_id}/rent", deps.rentalHandler.RentEquipment)
	router.HandleFunc("GET /equipments/{equipment_id}", deps.equipmentHandler.EquipmentById)
	router.HandleFunc("GET /equipments", deps.equipmentHandler.ListEquipmentHandler)
	router.HandleFunc("GET /users/{user_id}/equipments/lended", deps.equipmentHandler.GetEquipmentsByUserIdHandler)
	router.HandleFunc("PUT /user/{user_id}/equipments/{equipment_id}", deps.equipmentHandler.UpdateEquipmentHandler)
	router.HandleFunc("DELETE /user/{user_id}/equipments/{equipment_id}", deps.equipmentHandler.DeleteEquipmentHandler)

	return enableCORS(router)
}

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
