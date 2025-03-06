package equipment

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/middleware"
)

type equipmentHandler struct {
	eqipmentService Service
}

type Handler interface {
	CreateEquipmentHandler(w http.ResponseWriter, r *http.Request)
	ListEquipmentHandler(w http.ResponseWriter, r *http.Request)
	GetEquipmentsByUserIdHandler(w http.ResponseWriter, r *http.Request)
	DeleteEquipmentHandler(w http.ResponseWriter, r *http.Request)
	UpdateEquipmentHandler(w http.ResponseWriter, r *http.Request)
	EquipmentById(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &equipmentHandler{eqipmentService: service}
}

func (e *equipmentHandler) CreateEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	var equipment models.Equipment

	err := json.NewDecoder(r.Body).Decode(&equipment)

	if err != nil {
		utils.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidReqBody)
		return
	}

	resp, err := e.eqipmentService.CreateEquipment(context.Background(), equipment)
	if err != nil {
		if err == apperrors.ErrInvalidQuantity {
			utils.ErrorResponse(ctx, w, http.StatusBadRequest, err)
			return
		}
		utils.ErrorResponse(ctx, w, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(ctx, w, http.StatusCreated, resp)
}

func (e *equipmentHandler) ListEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	equipments, err := e.eqipmentService.GetAllEquipment(context.Background())

	if err != nil {
		log.Printf("Handler: error while fetching data from the backend: %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, equipments)
}

func (e *equipmentHandler) GetEquipmentsByUserIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("user_id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Handler: error while converting userId string into int: err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrPathParam)
		return
	}

	equipments, err := e.eqipmentService.GetEquipmentsByUserId(context.Background(), userId)

	if err != nil {
		switch err {
		case apperrors.ErrInvalidUserID:
			log.Printf("Handler: Invalid userId %d", userId)
			utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, err)
			return

		case apperrors.ErrNoData:
			log.Printf("Handler: No equipment found for userId %d", userId)
			utils.ErrorResponse(context.Background(), w, http.StatusNotFound, err)
			return

		default:
			log.Printf("Handler: Error fetching equipment for userId %d: %v", userId, err)
			utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, equipments)
}

func (e *equipmentHandler) DeleteEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	// token verification
	err := middleware.VerifyIncomingRequest(w, r)
	if err != nil {
		return
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	equipmentId, err := strconv.Atoi(r.PathValue("equipment_id"))
	if err != nil {
		log.Printf("Handler: error while converting EquipId req param values form string into int: %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
		return
	}

	// userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		log.Printf("Handler: error while converting UserId req param values form string into int: %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
		return
	}

	err = e.eqipmentService.DeleteEquipmentById(context.Background(), equipmentId, userId)
	if err != nil {
		switch err {
		case apperrors.ErrEquipmentNotFound:
			log.Printf("Handler: Equipment not found for equipId %d", equipmentId)
			utils.ErrorResponse(context.Background(), w, http.StatusNotFound, err)
			return

		case apperrors.ErrNotAnOwner:
			log.Printf("Handler: User %d is not the owner of equipment %d", userId, equipmentId)
			utils.ErrorResponse(context.Background(), w, http.StatusForbidden, err)
			return

		case apperrors.ErrNoData:
			log.Printf("Handler: No equipment found to delete for equipId %d", equipmentId)
			utils.ErrorResponse(context.Background(), w, http.StatusNotFound, err)
			return

		default:
			log.Printf("Handler: Internal server error while deleting equipId %d: %v", equipmentId, err)
			utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, "Equipment Deleted Successfully")
}

func (e *equipmentHandler) UpdateEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	// token verification
	err := middleware.VerifyIncomingRequest(w, r)
	if err != nil {
		return
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	equipmentId, err := strconv.Atoi(r.PathValue("equipment_id"))
	if err != nil {
		log.Printf("Handler: error while converting equipmentId param form string into int, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
		return
	}

	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		log.Printf("Handler: error while converting userId param form string into int, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
		return
	}

	var equipment models.Equipment
	err = json.NewDecoder(r.Body).Decode(&equipment)
	if err != nil {
		log.Printf("Handler: error while parsing request body, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrInvalidReqBody)
		return
	}

	resp, err := e.eqipmentService.UpdateEquipment(context.Background(), equipmentId, userId, equipment)
	if err != nil {
		switch err {
		case apperrors.ErrEquipmentNotFound:
			log.Printf("Handler: Equipment with ID %d not found", equipmentId)
			utils.ErrorResponse(context.Background(), w, http.StatusNotFound, err)
			return

		case apperrors.ErrNotAnOwner:
			log.Printf("Handler: User ID %d is not the owner of Equipment ID %d", userId, equipmentId)
			utils.ErrorResponse(context.Background(), w, http.StatusForbidden, err)
			return

		default:
			log.Printf("Handler: Error updating Equipment ID %d, err: %v", equipmentId, err)
			utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, resp)
}

func (e *equipmentHandler) EquipmentById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("equipment_id")

	equipId, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("Handler: error while converting equipment id param form string into int, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
		return
	}

	resp, err := e.eqipmentService.EquipmentById(context.Background(), equipId)
	if err != nil {
		switch err {
		case apperrors.ErrEquipmentNotFound:
			log.Printf("Handler: No equipment found with ID %d", equipId)
			utils.ErrorResponse(context.Background(), w, http.StatusNotFound, err)
			return

		default:
			log.Printf("Handler: Error fetching equipment with ID %d: %v", equipId, err)
			utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, resp)
}
