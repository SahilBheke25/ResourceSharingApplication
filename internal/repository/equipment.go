package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
)

const (
	createEquipment = `INSERT INTO equipments (
						  				equipment_name, description, rent_per_day, quantity, equipment_img, user_id) 
											VALUES ($1, $2, $3, $4, $5, $6) 
											RETURNING id, equipment_name, description, rent_per_day, quantity, 
											equipment_img, status, uploaded_at`

	getEquipments = `SELECT id, equipment_name, description, rent_per_day, quantity, 
						  			equipment_img, status, uploaded_at from equipments`

	equipmentsByUserId = `SELECT id, equipment_name, description, rent_per_day, 
												quantity, equipment_img, status, uploaded_at from equipments 
												WHERE user_id = $1`

	deleteEquipment = `DELETE FROM equipments WHERE id = $1`

	updateEquipment = `UPDATE equipments SET equipment_name = $1,
    									description = $2,
    									rent_per_day = $3,
    									quantity = $4,
   										status = $5,
    									equipment_img = $6
											WHERE id = $7
											RETURNING id, equipment_name, description, rent_per_day, quantity, 
											equipment_img, status, uploaded_at`

	getEquipmentQuantity = `SELECT quantity from equipments WHERE id = $1`

	getEquipmentCharges = `SELECT rent_per_day from equipments WHERE id = $1`

	descreaseQuantity = `UPDATE equipments SET quantity = quantity - $1
												WHERE id = $2`

	equipmentById = `SELECT id, equipment_name, description, rent_per_day, quantity, equipment_img, status, uploaded_at, user_id FROM equipments WHERE id = $1`
)

type equipment struct {
	db *sql.DB
}

type EquipmentStorer interface {
	CreateEquipment(ctx context.Context, eqp models.Equipment) (models.Equipment, error)
	GetAllEquipment(ctx context.Context) ([]models.Equipment, error)
	GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error)
	DeleteEquipmentById(ctx context.Context, equipId int) error
	UpdateEquipment(tx context.Context, equipId int, equipment models.Equipment) (models.Equipment, error)
	EquipmentById(ctx context.Context, equipId int) (models.Equipment, error)
}

func NewEquipmentStore(db *sql.DB) EquipmentStorer {
	return equipment{db: db}
}

func (e equipment) CreateEquipment(ctx context.Context, eqp models.Equipment) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, createEquipment,
		eqp.Name,
		eqp.Description,
		eqp.RentPerDay,
		eqp.Quantity,
		eqp.EquipmentImg,
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
		&resp.RentPerDay,
		&resp.Quantity,
		&resp.EquipmentImg,
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

		err := list.Scan(
			&equipment.ID,
			&equipment.Name,
			&equipment.Description,
			&equipment.RentPerDay,
			&equipment.Quantity,
			&equipment.EquipmentImg,
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
			&equipment.RentPerDay,
			&equipment.Quantity,
			&equipment.EquipmentImg,
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

func (e equipment) DeleteEquipmentById(ctx context.Context, equipId int) error {

	res, err := e.db.Exec(deleteEquipment, equipId)

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

func (e equipment) UpdateEquipment(ctx context.Context, equipId int, equipment models.Equipment) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, updateEquipment,
		equipment.Name,
		equipment.Description,
		equipment.RentPerDay,
		equipment.Quantity,
		equipment.Status,
		equipment.EquipmentImg,
		equipId)

	err := res.Err()

	if err != nil {
		return models.Equipment{}, fmt.Errorf("error while Deleting equipment: %v", err)
	}

	var resp models.Equipment

	err = res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerDay,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.Status,
		&resp.UploadedAt)

	if err != nil {
		log.Println("error while scaning query result in update equipment, err : ", err)
		return models.Equipment{}, apperrors.ErrDbScan
	}

	return resp, nil
}

func (e equipment) EquipmentById(ctx context.Context, equipId int) (models.Equipment, error) {

	res := e.db.QueryRow(equipmentById, equipId)

	var resp models.Equipment

	err := res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerDay,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.Status,
		&resp.UploadedAt,
		&resp.UserId)

	if err != nil {
		log.Println("error while scaning query result in equipmentById, err : ", err)
		return models.Equipment{}, apperrors.ErrDbScan
	}

	return resp, nil
}
