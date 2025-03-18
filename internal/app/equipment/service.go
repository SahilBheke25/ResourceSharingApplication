package equipment

import (
	"context"
	"log"

	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
	"github.com/SahilBheke25/quick-farm-backend/internal/repository"
)

type service struct {
	equipmentRepo repository.EquipmentStorer
}

type Service interface {
	CreateEquipment(ctx context.Context, equipment models.Equipment) (models.Equipment, error)
	GetAllEquipment(ctx context.Context) ([]models.Equipment, error)
	GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error)
	DeleteEquipmentById(ctx context.Context, equipmentId int, userId int) error
	UpdateEquipment(ctx context.Context, equipmentId int, userId int, equipment models.Equipment) (models.Equipment, error)
	EquipmentById(ctx context.Context, equipId int) (models.Equipment, error)
}

// constructor function to initialize service layer dependency for equipments
func NewService(eqp repository.EquipmentStorer) Service {
	return service{equipmentRepo: eqp}
}

func (s service) CreateEquipment(ctx context.Context, equipment models.Equipment) (models.Equipment, error) {

	if equipment.Quantity <= 0 {
		log.Println("error: equipment quantity must be greater than zero")
		return models.Equipment{}, apperrors.ErrInvalidQuantity
	}

	resp, err := s.equipmentRepo.CreateEquipment(ctx, equipment)
	if err != nil {
		log.Printf("error occured while calling CreateEquipemnt DB opeartion, err : %v", err)
		return models.Equipment{}, err
	}

	return resp, nil
}

func (s service) GetAllEquipment(ctx context.Context) ([]models.Equipment, error) {
	resp, err := s.equipmentRepo.GetAllEquipment(ctx)

	if err != nil {
		log.Printf("Service: error occured while calling CreateEquipemnt DB opeartion, err : %v", err)
		return []models.Equipment{}, err
	}

	return resp, nil
}

func (s service) GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error) {

	if userId <= 0 {
		log.Printf("Service: error invalid userId:%v provided, err : %v", userId, apperrors.ErrInvalidUserID)
		return nil, apperrors.ErrInvalidUserID
	}

	resp, err := s.equipmentRepo.EquipmentsOfUser(ctx, userId)
	if err != nil {
		log.Printf("Service: error occured while calling get Equipment by user ID DB opeartion, err : %v", err)
		return nil, err
	}

	return resp, nil
}

func (s service) DeleteEquipmentById(ctx context.Context, equipmentId int, userId int) error {

	equipment, err := s.equipmentRepo.EquipmentById(ctx, equipmentId)
	if err != nil {
		log.Printf("Service: error while calling EquipmentById DB operation, err : %v", err)
		return apperrors.ErrEquipmentNotFound
	}

	if equipment.UserId != userId {
		log.Printf("Service: error user with userId:%v is not an owner of the equipment:%v", userId, equipmentId)
		return apperrors.ErrNotAnOwner
	}

	err = s.equipmentRepo.DeleteEquipmentById(ctx, equipmentId)
	if err != nil {
		log.Printf("Service: error while calling DeleteEquipmentById DB operation, err : %v", err)
		return err
	}

	return nil
}

func (s service) UpdateEquipment(ctx context.Context, equipmentId int, userId int, equipment models.Equipment) (models.Equipment, error) {

	currEquipment, err := s.equipmentRepo.EquipmentById(ctx, equipmentId)
	if err != nil {
		log.Printf("Service: error while calling EquipmentById DB operation, err : %v", err)
		return models.Equipment{}, apperrors.ErrEquipmentNotFound
	}

	if currEquipment.UserId != userId {
		log.Printf("Service: error user with userId:%v is not an owner of the equipment:%v", userId, equipmentId)
		return models.Equipment{}, apperrors.ErrNotAnOwner
	}

	resp, err := s.equipmentRepo.UpdateEquipment(ctx, equipmentId, userId, equipment)
	if err != nil {
		log.Printf("Service: error occurred while updating Equipment ID %d, err: %v", equipmentId, err)
		return models.Equipment{}, err
	}

	return resp, nil
}

func (s service) EquipmentById(ctx context.Context, equipId int) (models.Equipment, error) {

	resp, err := s.equipmentRepo.EquipmentById(ctx, equipId)
	if err != nil {
		log.Printf("Service: error occured during calling EquipmentById DB operation, err : %v", err)
		return models.Equipment{}, err
	}

	return resp, nil
}
