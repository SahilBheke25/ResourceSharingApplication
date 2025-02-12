package rental

import (
	"context"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

type service struct {
	rentalRepo repository.RentalStorer
}

type Service interface {
	RentEquipment(ctx context.Context, rental models.Rental) (models.Billing, error)
}

func NewService(rentalRepo repository.RentalStorer) Service {

	return service{rentalRepo: rentalRepo}
}

func (s service) RentEquipment(ctx context.Context, rental models.Rental) (models.Billing, error) {

	rental.Duration = (rental.RentTill.Sub(rental.RentAt)).Hours()

	if rental.Duration < 0.5 {
		return models.Billing{}, apperrors.ErrDurationTooShort
	}

	availbaleQuntity, err := s.rentalRepo.EquipmentQuantity(ctx, rental.EquipId)

	if err != nil {
		return models.Billing{}, err
	}

	if availbaleQuntity < rental.Quantity {
		return models.Billing{}, apperrors.ErrQuantityNotAvailable
	}

	rentPerHour, err := s.rentalRepo.EquipmentCharges(ctx, rental.EquipId)

	if err != nil {
		return models.Billing{}, err
	}

	resp, err := s.rentalRepo.RentEquipment(ctx, rental)

	if err != nil {
		return models.Billing{}, err
	}

	err = s.rentalRepo.UpdateQuantity(ctx, rental.EquipId, rental.Quantity)
	if err != nil {
		return models.Billing{}, err
	}

	duration := resp.Duration

	var billing models.Billing

	billing.Amount = duration * float64(rentPerHour)
	billing.RentId = resp.Id

	bill, _ := s.rentalRepo.CreateBill(ctx, billing)

	return bill, nil
}
