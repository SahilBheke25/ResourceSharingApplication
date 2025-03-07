package models

import "time"

type Rental struct {
	Id       int       `json:"id,omitempty"`
	RentAt   time.Time `json:"rent_at"`
	RentTill time.Time `json:"rent_till"`
	Duration float64   `json:"duration"`
	Quantity int       `json:"quantity"`
	EquipId  int       `json:"equipment_id"`
	UserId   int       `json:"user_id"`
}

// type RentedEquipmentData struct {
// 	Id       int       `json:"id,omitempty"`
// 	RentAt   time.Time `json:"rent_at"`
// 	RentTill time.Time `json:"rent_till"`
// 	Duration float64   `json:"duration"`
// 	Quantity int       `json:"quantity"`
// 	EquipId  int       `json:"equipment_id"`
// 	UserId   int       `json:"user_id"`
// 	Amount   float64   `json:"amount"`
// }
