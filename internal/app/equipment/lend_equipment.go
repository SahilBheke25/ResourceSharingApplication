package Equipment

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func GetAllEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	equipments, err := Repository.GetAllEquipment()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	handle.HandleResponse(w, equipments, r)
}
