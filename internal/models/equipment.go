package models

import "time"

type Equipment struct {
	ID            string    `json:"id,omitempty"`
	Name          string    `json:"equipment_name"`
	Description   string    `json:"description"`
	RentPerHour   int       `json:"rent_per_hour"`
	Quantity      int       `json:"quantity"`
	EquipmentImg  string    `json:"equipment_img"`
	AvailableFrom time.Time `json:"available_from"`
	AvailableTill time.Time `json:"available_till"`
	UserId        int       `json:"user_id,omitempty"`
	Status        string    `json:"status"`
	UploadedAt    string    `json:"uploaded_at"`
}
