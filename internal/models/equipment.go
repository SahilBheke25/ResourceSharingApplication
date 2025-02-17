package models

type Equipment struct {
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"equipment_name"`
	Description  string  `json:"description"`
	RentPerDay   float64 `json:"rent_per_day"`
	Quantity     int     `json:"quantity"`
	EquipmentImg string  `json:"equipment_img"`
	UserId       int     `json:"user_id,omitempty"`
	Status       string  `json:"status"`
	UploadedAt   string  `json:"uploaded_at"`
}
