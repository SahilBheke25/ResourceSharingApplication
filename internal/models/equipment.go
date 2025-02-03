package Models

type Equipment struct {
	EquipmentName string `json:"equipmentname"`
	Description   string `json:"description"`
	RentPerHour   int    `json:"rentperhour"`
	Quantity      int    `json:"quantity"`
	EquipmentImg  string `json:"equipmentimage"`
	AvailableFrom string `json: "availablefrom"`
	AvailableTill string `json: "availabletill:`
	Username      string `json: :username"`
}
