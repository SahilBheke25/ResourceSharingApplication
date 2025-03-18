package rental

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/quick-farm-backend/internal/app/utils"
	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
)

type rentalHandler struct {
	rentalService Service
}

type Handler interface {
	RentEquipment(w http.ResponseWriter, r *http.Request)
}

func NewHandler(rentalService Service) Handler {
	return &rentalHandler{rentalService: rentalService}
}

func (rentalH *rentalHandler) RentEquipment(w http.ResponseWriter, r *http.Request) {

	// Path param conversion
	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		log.Printf("Handler: error while converting user id param form string to int, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
	}

	// Path param conversion
	equipId, err := strconv.Atoi(r.PathValue("equip_id"))
	if err != nil {
		log.Printf("Handler: error while converting equipment id param form string to int, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrAtoi)
	}

	// Reading json request
	var rental models.Rental
	err = json.NewDecoder(r.Body).Decode(&rental)
	if err != nil {
		log.Printf("Handler: error in RentEquipment handler while parsing request body, err : %v", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrInvalidReqBody)
		return
	}

	// Calling renting service
	rental.UserId, rental.EquipId = userId, equipId
	resp, err := rentalH.rentalService.RentEquipment(context.Background(), rental)
	if err != nil {
		log.Println("error in RentEquipment handler while making request to service layer, err : ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(context.Background(), w, http.StatusCreated, resp)
}
