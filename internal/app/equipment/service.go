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
	DeleteEquipmentById(ctx context.Context, equipmentId int) error
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

func (s service) DeleteEquipmentById(ctx context.Context, equipmentId int) error {

	err := s.equipmentRepo.DeleteEquipmentById(ctx, equipmentId)

	if err != nil {
		log.Println("eerror while calling DeleteEquipmentById DB operation, err : ", err)
		return err
	}

	return nil
}
