package rental

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
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
		log.Println("error while converting equipment id param form string into int, err : ", err)
		http.Error(w, "error while paring path value", http.StatusInternalServerError)
	}

	// Path param conversion
	equipId, err := strconv.Atoi(r.PathValue("equip_id"))
	if err != nil {
		log.Println("error while converting equipment id param form string into int, err : ", err)
		http.Error(w, "error while path value type casting", http.StatusInternalServerError)
	}

	// Reading json request
	var rental models.Rental
	err = json.NewDecoder(r.Body).Decode(&rental)
	if err != nil {
		log.Println("error in RentEquipment handler while parsing request body, err : ", err)
		http.Error(w, "Error while Decoding Rental Equipment Request Body", http.StatusBadRequest)
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

	utils.HandleResponse(w, resp, r)
}
