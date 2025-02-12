package models

import "time"

type Billing struct {
	Id     int       `json:"id"`
	Date   time.Time `json:"payment_date"`
	Amount float64   `json:"total_amount"`
	Type   string    `json:"payment_type,omitempty"`
	Status string    `json:"payment_status,omitempty"`
	RentId int       `json:"rent_id"`
}
