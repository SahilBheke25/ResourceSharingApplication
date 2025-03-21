package rental

import (
	"context"
	"log"

	"github.com/SahilBheke25/quick-farm-backend/internal/app/equipment"
	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
	"github.com/SahilBheke25/quick-farm-backend/internal/repository"
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

	// Duration validation
	rental.Duration = (rental.RentTill.Sub(rental.RentAt)).Hours()
	if rental.Duration < 24 {
		log.Printf("Service: error duration must be atleast 1 Day")
		err = apperrors.ErrDurationTooShort
		return
	}

	// Equipment to rent
	equip, err := s.equipmentService.EquipmentById(ctx, rental.EquipId)
	if err != nil {
		log.Printf("Service: error while calling equipment Service from rent equipment service, err : %v", err)
		return models.Billing{}, err
	}

	// Quantity check
	availbaleQuntity := equip.Quantity
	if availbaleQuntity < rental.Quantity {
		log.Printf("Service: error quantity not available")
		return models.Billing{}, apperrors.ErrQuantityNotAvailable
	}

	// Renting DB call
	rentPerDay := equip.RentPerDay
	bill, err = s.rentalRepo.RentEquipment(ctx, rental, availbaleQuntity, rentPerDay)
	if err != nil {
		log.Printf("Service: error while calling RentEquipment DB operation, err : %v", err)
		return models.Billing{}, err
	}

	return bill, nil
}
