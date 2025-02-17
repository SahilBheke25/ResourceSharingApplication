package rental

import (
	"context"
	"log"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

type service struct {
	rentalRepo       repository.RentalStorer
	equipmentService equipment.Service
}

type Service interface {
	RentEquipment(ctx context.Context, rental models.Rental) (models.Billing, error)
}

func NewService(rentalRepo repository.RentalStorer, equipmentService equipment.Service) Service {
	return service{rentalRepo: rentalRepo, equipmentService: equipmentService}
}

func (s service) RentEquipment(ctx context.Context, rental models.Rental) (bill models.Billing, err error) {

	rental.Duration = (rental.RentTill.Sub(rental.RentAt)).Hours()
	if rental.Duration < 0.5 {
		err = apperrors.ErrDurationTooShort
		return
	}

	equip, err := s.equipmentService.EquipmentById(ctx, rental.EquipId)
	if err != nil {
		log.Println("error while calling equipment Service from rent equipment service, err : ", err)
		return models.Billing{}, err
	}

	availbaleQuntity := equip.Quantity
	if availbaleQuntity < rental.Quantity {
		return models.Billing{}, apperrors.ErrQuantityNotAvailable
	}

	rentPerDay := equip.RentPerDay
	bill, err = s.rentalRepo.RentEquipment(ctx, rental, availbaleQuntity, rentPerDay)
	if err != nil {
		return models.Billing{}, err
	}

	return bill, nil
}
