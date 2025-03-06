package repository

import (
	"context"
	"database/sql"
	"errors"
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
	EquipmentsOfUser(ctx context.Context, userId int) ([]models.Equipment, error)
	DeleteEquipmentById(ctx context.Context, equipId int) error
	UpdateEquipment(tx context.Context, equipId int, userId int, equipment models.Equipment) (models.Equipment, error)
	EquipmentById(ctx context.Context, equipId int) (models.Equipment, error)
}

func NewEquipmentStore(db *sql.DB) EquipmentStorer {
	return equipment{db: db}
}

func (e equipment) CreateEquipment(ctx context.Context, eqp models.Equipment) (models.Equipment, error) {

	var resp models.Equipment

	res := e.db.QueryRowContext(ctx, createEquipment,
		eqp.Name,
		eqp.Description,
		eqp.RentPerDay,
		eqp.Quantity,
		eqp.EquipmentImg,
		eqp.UserId,
	)

	err := res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerDay,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.Status,
		&resp.UploadedAt)

	if err != nil {
		log.Printf("Repo: Database error while creating equipment, err : %v", err)
		return models.Equipment{}, apperrors.ErrFailedToCreate
	}

	return resp, nil
}

func (e equipment) GetAllEquipment(ctx context.Context) ([]models.Equipment, error) {

	var equipmentArr []models.Equipment

	list, err := e.db.Query(getEquipments)

	if err != nil {
		log.Printf("Repo: Error while executing query, err : %v", err)
		return equipmentArr, apperrors.ErrDbFetching
	}
	defer list.Close()

	for list.Next() {
		var equipment models.Equipment

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
			log.Printf("Repo: Error while scanning row , err : %v", err)
			return equipmentArr, apperrors.ErrDbScan
		}
		equipmentArr = append(equipmentArr, equipment)
	}

	return equipmentArr, nil
}

func (e equipment) EquipmentsOfUser(ctx context.Context, userId int) ([]models.Equipment, error) {

	var equipmentArr []models.Equipment

	list, err := e.db.Query(equipmentsByUserId, userId)
	if err != nil {
		log.Printf("Repo: error while executing query, err : %v", err)
		return equipmentArr, apperrors.ErrDbFetching
	}
	defer list.Close()

	hasRows := false

	for list.Next() {
		hasRows = true

		var equipment models.Equipment
		err := list.Scan(&equipment.ID,
			&equipment.Name,
			&equipment.Description,
			&equipment.RentPerDay,
			&equipment.Quantity,
			&equipment.EquipmentImg,
			&equipment.Status,
			&equipment.UploadedAt)

		if err != nil {
			log.Printf("Repo: error while accessing DB, err : %v", err)
			return equipmentArr, apperrors.ErrDbScan
		}

		equipmentArr = append(equipmentArr, equipment)
	}

	if err = list.Err(); err != nil {
		log.Printf("Repo: error after iterating over rows, err : %v", err)
		return equipmentArr, apperrors.ErrDbServer
	}

	if !hasRows {
		return equipmentArr, apperrors.ErrNoData
	}

	return equipmentArr, nil
}

func (e equipment) DeleteEquipmentById(ctx context.Context, equipId int) error {

	res, err := e.db.Exec(deleteEquipment, equipId)
	if err != nil {
		log.Printf("Repo: error while deleting equipment (equipId: %d), err: %v", equipId, err)
		return apperrors.ErrDbDelete
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Repo: error fetching RowsAffected for equipId %d, err: %v", equipId, err)
		return apperrors.ErrDbServer
	}
	if count == 0 {
		log.Printf("Repo: No equipment found to delete for equipId %d", equipId)
		return apperrors.ErrNoData
	}

	return nil
}

func (e equipment) UpdateEquipment(ctx context.Context, equipId int, userId int, equipment models.Equipment) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, updateEquipment,
		equipment.Name,
		equipment.Description,
		equipment.RentPerDay,
		equipment.Quantity,
		equipment.Status,
		equipment.EquipmentImg,
		equipId)

	var resp models.Equipment
	err := res.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.RentPerDay,
		&resp.Quantity,
		&resp.EquipmentImg,
		&resp.Status,
		&resp.UploadedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Repo: No equipment found with ID %d for update", equipId)
			return models.Equipment{}, apperrors.ErrNoData
		}

		log.Printf("Repo: Error while scanning query result in UpdateEquipment, err: %v", err)
		return models.Equipment{}, apperrors.ErrDbScan
	}

	log.Println("INFO Repo: Equipment updated successfully: ", resp)

	return resp, nil
}

func (e equipment) EquipmentById(ctx context.Context, equipId int) (models.Equipment, error) {

	res := e.db.QueryRowContext(ctx, equipmentById, equipId)

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
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Repo: No equipment found with ID %d", equipId)
			return models.Equipment{}, apperrors.ErrEquipmentNotFound
		}

		log.Printf("Repo: Error scanning query result in EquipmentById (ID: %d), err: %v", equipId, err)
		return models.Equipment{}, apperrors.ErrDbScan
	}

	return resp, nil
}
