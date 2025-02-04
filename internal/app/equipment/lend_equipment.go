package Equipment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/handle"
	Models "github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	Repository "github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

func PostLendEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("IN PostLendEquipmentHandler")

	var equipment Models.Equipment
	err := json.NewDecoder(r.Body).Decode(&equipment)
	if err != nil {
		http.Error(w, "Error while Decoding Request Body", http.StatusBadRequest)
	}
	// equipment_name, description, equipment_image, available_from, available_till, username, string, rent_per_hour, quantity,
	err = Repository.CreateEquipment(equipment.EquipmentName, equipment.Description, equipment.EquipmentImg, equipment.AvailableFrom, equipment.AvailableTill, equipment.Username, equipment.RentPerHour, equipment.Quantity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handle.HandleResponse(w, "Equipment Listed Successfully ", r)

	fmt.Println(equipment)
	handle.HandleResponse(w, equipment, r)
}

func ListEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	equipments, err := Repository.GetAllEquipment()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	handle.HandleResponse(w, equipments, r)
}

func GetEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	var userID string

	// Extraction path param using trip fuction
	if strings.HasPrefix(path, "/equipments/") {
		userID = strings.TrimPrefix(path, "/equipments/") // Extract the part after "/equipments/"

		// Do something with the username, e.g., fetch equipment for this user
		fmt.Fprintf(w, "Equipment for user: %s", userID)
	} else {
		http.NotFound(w, r) // Return 404 if the path doesn't match
	}

	id, err := strconv.Atoi(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Id; ", id)
	// equipments, err := Repository.GetEquipment(id)

	handle.HandleResponse(w, "equipments", r)
}
