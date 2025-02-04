package Repository

import (
	"fmt"

	Models "github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

const (
	createNewEquipment = `INSERT INTO equipments (equipment_name, description, rent_per_hour, quantity, equipment_img, available_from, available_till, username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	getEquipments      = `SELECT equipment_name, description, rent_per_hour, quantity, equipment_img, available_from, available_till, username from equipments`
	quipmentByusername = `SELECT password from users where user_name = $1`
)

func CreateEquipment(equipment_name, description, equipment_img, available_from, available_till, username string, rent_per_hour, quantity int) error {

	_, err := DB.Exec(createNewEquipment, equipment_name, description, rent_per_hour, quantity, equipment_img, available_from, available_till, username)

	if err != nil {
		return fmt.Errorf("Error while creating new Equipment entry: %v", err)
	}

	return nil
}

func GetAllEquipment() ([]Models.Equipment, error) {

	var equipment Models.Equipment
	var equipmentArr []Models.Equipment

	list, err := DB.Query(getEquipments)
	if err != nil {
		err = fmt.Errorf("Error while executing query: %v", err)
		return equipmentArr, err
	}

	for list.Next() {
		err := list.Scan(&equipment.EquipmentName, &equipment.Description, &equipment.RentPerHour, &equipment.Quantity, &equipment.EquipmentImg, &equipment.AvailableFrom, &equipment.AvailableTill, &equipment.Username)
		if err != nil {
			err = fmt.Errorf("Error while accessing DB: %v", err)
			return equipmentArr, err
		}
		fmt.Println(equipment)
		equipmentArr = append(equipmentArr, equipment)

	}
	return equipmentArr, nil
}

func GetEquipment(username string) {

}

// func AuthenticateUser(userName, password string) (bool, error) {

// 	var dbPassword string

// 	err := DB.QueryRow(userByusername, userName).Scan(&dbPassword)

// 	if err != nil {
// 		return false, fmt.Errorf("User Not Found: %v", err)
// 	}

// 	if dbPassword != password {
// 		return false, fmt.Errorf("Wrong Password!!")
// 	}

// 	return true, nil
// }
