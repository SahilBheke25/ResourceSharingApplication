package equipment

import (
	"context"
	"log"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

type service struct {
	equipmentRepo repository.EquipmentStorer
}

type Service interface {
	CreateEquipment(ctx context.Context, equipment models.Equipment) (models.Equipment, error)
	GetAllEquipment(ctx context.Context) ([]models.Equipment, error)
	GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error)
	DeleteEquipmentById(ctx context.Context, equipmentId int) error
	UpdateEquipment(ctx context.Context, equipmentId int, equipment models.Equipment) (models.Equipment, error)
	EquipmentById(ctx context.Context, equipId int) (models.Equipment, error)
}

// constructor function to initialize service layer dependency for equipments
func NewService(eqp repository.EquipmentStorer) Service {
	return service{equipmentRepo: eqp}
}

func (s service) CreateEquipment(ctx context.Context, equipment models.Equipment) (models.Equipment, error) {
	resp, err := s.equipmentRepo.CreateEquipment(ctx, equipment)
	if err != nil {
		log.Println("error occured while calling CreateEquipemnt DB opeartion, err : ", err)
		return models.Equipment{}, err
	}

	return resp, nil
}

func (s service) GetAllEquipment(ctx context.Context) ([]models.Equipment, error) {
	resp, err := s.equipmentRepo.GetAllEquipment(ctx)

	if err != nil {
		log.Println("error occured while calling CreateEquipemnt DB opeartion, err : ", err)
		return []models.Equipment{}, err
	}

	return resp, nil

}

func (s service) GetEquipmentsByUserId(ctx context.Context, userId int) ([]models.Equipment, error) {
	resp, err := s.equipmentRepo.GetEquipmentsByUserId(ctx, userId)

	if err != nil {
		log.Println("error occured while calling get Equipment by user ID DB opeartion, err : ", err)
		return []models.Equipment{}, err
	}

	return resp, nil
}

func (s service) DeleteEquipmentById(ctx context.Context, equipmentId int) error {

	err := s.equipmentRepo.DeleteEquipmentById(ctx, equipmentId)

	if err != nil {
		log.Println("eerror while calling DeleteEquipmentById DB operation, err : ", err)
		return err
	}

	return nil
}

func (s service) UpdateEquipment(ctx context.Context, equipmentId int, equipment models.Equipment) (models.Equipment, error) {

	resp, err := s.equipmentRepo.UpdateEquipment(ctx, equipmentId, equipment)

	if err != nil {
		log.Println("error occured while calling Update Equipment DB opeartion, err : ", err)
		return models.Equipment{}, err
	}

	return resp, nil
}

func (s service) EquipmentById(ctx context.Context, equipId int) (models.Equipment, error) {

	resp, err := s.equipmentRepo.EquipmentById(ctx, equipId)
	if err != nil {
		log.Panicln("error occured during calling EquipmentById DB operation, err : ", err)
		return models.Equipment{}, err
	}

	return resp, nil
}
