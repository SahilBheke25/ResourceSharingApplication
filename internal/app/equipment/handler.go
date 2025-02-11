package equipment

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

type equipmentHandler struct {
	eqipmentService Service
}

type Handler interface {
	CreateEquipmentHandler(w http.ResponseWriter, r *http.Request)
	ListEquipmentHandler(w http.ResponseWriter, r *http.Request)
	GetEquipmentsByUserIdHandler(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &equipmentHandler{eqipmentService: service}
}

func (e *equipmentHandler) CreateEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	var equipment models.Equipment

	err := json.NewDecoder(r.Body).Decode(&equipment)

	if err != nil {

		http.Error(w, "Error while Decoding Request Body", http.StatusBadRequest)
	}

	resp, err := e.eqipmentService.CreateEquipment(context.Background(), equipment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HandleResponse(w, resp, r)
}

func (e *equipmentHandler) ListEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	equipments, err := e.eqipmentService.GetAllEquipment(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	utils.HandleResponse(w, equipments, r)
}

func (e *equipmentHandler) GetEquipmentsByUserIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("user_id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("error while converting userId string into int: %v", err)
		http.Error(w, "Invalid userId in path value", http.StatusBadRequest)
		return
	}

	equipments, err := e.eqipmentService.GetEquipmentsByUserId(context.Background(), userId)

	if err != nil {
		log.Printf("error while fetching data from the backend: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(equipments) == 0 {
		utils.HandleResponse(w, "Data Not Found", r)
		return
	}

	utils.HandleResponse(w, equipments, r)
}
