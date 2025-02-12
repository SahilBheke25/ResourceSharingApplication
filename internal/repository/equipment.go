package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

const (
	createEquipment = `INSERT INTO equipments (
						  	equipment_name, description, rent_per_hour, quantity, 
							equipment_img, available_from, available_till, user_id) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
							RETURNING id, equipment_name, description, rent_per_hour, quantity, 
							equipment_img, available_from, available_till, status, uploaded_at`

	getEquipments = `SELECT equipment_name, description, rent_per_hour, quantity, 
						  	equipment_img, available_from, available_till, status, 
							uploaded_at from equipments`

	equipmentsByUserId = `SELECT id, equipment_name, description, rent_per_hour, 
							quantity, equipment_img, available_from, available_till, 
							status, uploaded_at from equipments 
							WHERE user_id = $1`

	deleteEquipment = `DELETE FROM equipments WHERE id = $1`

	updateEquipment = `UPDATE equipments SET equipment_name = $1,
    					description = $2,
    					rent_per_hour = $3,
    					quantity = $4,
   						status = $5,
    					equipment_img = $6,
    					available_from = $7,
                        available_till = $8
						WHERE id = $9
						RETURNING id, equipment_name, description, rent_per_hour, quantity, 
						equipment_img, available_from, available_till, status, uploaded_at`
)

type equipment struct {
	db *sql.DB
}

type EquipmentStorer interface {
	CreateEquipment(ctx context.Context, eqp models.Equipment) (models.Equipment, error)
	GetAllEquipment(ctx context.Context) ([]models.Equipment, error)
	GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error)
	DeleteEquipmentById(ctx context.Context, equipmentId int) error
	UpdateEquipment(tx context.Context, equipmentId int, equipment models.Equipment) (models.Equipment, error)
}

func NewEquipmentStore(db *sql.DB) EquipmentStorer {
	return equipment{db: db}
}

func (e equipment) CreateEquipment(ctx context.Context, eqp models.Equipment) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, createEquipment,
		eqp.Name,
		eqp.Description,
		eqp.RentPerHour,
		eqp.Quantity,
		eqp.EquipmentImg,
		eqp.AvailableFrom,
		eqp.AvailableTill,
		eqp.UserId,
	)
	err := res.Err()
	if err != nil {
		log.Println("error occured while making db reqeust for create quipment, err : ", err)
		return models.Equipment{}, err
	}

	var resp models.Equipment

	res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerHour,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.AvailableFrom,
		&resp.AvailableTill,
		&resp.Status,
		&resp.UploadedAt)

	return resp, nil
}

func (e equipment) GetAllEquipment(ctx context.Context) ([]models.Equipment, error) {

	var equipment models.Equipment
	var equipmentArr []models.Equipment

	list, err := e.db.Query(getEquipments)

	if err != nil {
		log.Println("error while executing query, err : ", err)
		return equipmentArr, err
	}

	for list.Next() {

		err := list.Scan(&equipment.Name,
			&equipment.Description,
			&equipment.RentPerHour,
			&equipment.Quantity,
			&equipment.EquipmentImg,
			&equipment.AvailableFrom,
			&equipment.AvailableTill,
			&equipment.Status,
			&equipment.UploadedAt)

		if err != nil {
			log.Println("error while accessing DB, err : ", err)
			return equipmentArr, err
		}

		equipmentArr = append(equipmentArr, equipment)
	}

	return equipmentArr, nil
}

func (e equipment) GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error) {

	var equipment models.Equipment
	var equipmentArr []models.Equipment

	list, err := e.db.Query(equipmentsByUserId, userId)

	if err != nil {
		log.Println("error while executing query, err : ", err)
		return equipmentArr, err
	}

	for list.Next() {

		err := list.Scan(&equipment.ID,
			&equipment.Name,
			&equipment.Description,
			&equipment.RentPerHour,
			&equipment.Quantity,
			&equipment.EquipmentImg,
			&equipment.AvailableFrom,
			&equipment.AvailableTill,
			&equipment.Status,
			&equipment.UploadedAt)

		if err != nil {
			log.Println("error while accessing DB, err : ", err)
			return equipmentArr, err
		}

		equipmentArr = append(equipmentArr, equipment)
	}

	return equipmentArr, nil
}

func (e equipment) DeleteEquipmentById(ctx context.Context, equipmentId int) error {

	res, err := e.db.Exec(deleteEquipment, equipmentId)

	if err != nil {
		log.Println("error while Deleting equipment, err : ", err)
		return err
	}

	var count, _ = res.RowsAffected()

	if count == 0 {
		return fmt.Errorf("no data found Bad Request")
	}

	return nil
}

func (e equipment) UpdateEquipment(ctx context.Context, equipmentId int, equipment models.Equipment) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, updateEquipment,
		equipment.Name,
		equipment.Description,
		equipment.RentPerHour,
		equipment.Quantity,
		equipment.Status,
		equipment.EquipmentImg,
		equipment.AvailableFrom,
		equipment.AvailableTill,
		equipmentId)

	err := res.Err()

	if err != nil {
		return models.Equipment{}, fmt.Errorf("error while Deleting equipment: %v", err)
	}

	var resp models.Equipment

	res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerHour,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.AvailableFrom,
		&resp.AvailableTill,
		&resp.Status,
		&resp.UploadedAt)

	return resp, nil
}
