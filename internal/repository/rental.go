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
	createRental = `INSERT INTO rental(quantity, rent_at, rent_till, duration, equip_id, user_id) 
									VALUES($1, $2, $3, $4, $5, $6)
									RETURNING id, quantity, rent_at, rent_till, duration, equip_id, user_id`

	createNewBill = `INSERT	INTO billing(total_amount, rent_id) VALUES($1, $2)
										RETURNING id, payment_date, total_amount, rent_id`
)

type Rental struct {
	db *sql.DB
}

type RentalStorer interface {
	RentEquipment(ctx context.Context, rental models.Rental) (models.Rental, error)
	EquipmentQuantity(ctx context.Context, equipId int) (int, error)
	EquipmentCharges(ctx context.Context, equipId int) (int, error)
	CreateBill(ctx context.Context, billing models.Billing) (models.Billing, error)
	UpdateQuantity(ctx context.Context, equip_id int, quantity int) error
}

func NewRentalStore(db *sql.DB) RentalStorer {
	return Rental{db: db}
}

func (r Rental) RentEquipment(ctx context.Context, rental models.Rental) (models.Rental, error) {

	res := r.db.QueryRowContext(ctx, createRental,
		rental.Quantity,
		rental.RentAt,
		rental.RentTill,
		rental.Duration,
		rental.EquipId,
		rental.UserId)

	err := res.Err()

	if err != nil {
		log.Println("error occured while making db reqeust for create reantal, err : ", err)
		return models.Rental{}, err
	}

	var resp models.Rental

	res.Scan(
		&resp.Id,
		&resp.Quantity,
		&resp.RentAt,
		&resp.RentTill,
		&resp.Duration,
		&resp.EquipId,
		&resp.UserId)

	return resp, nil
}

func (r Rental) EquipmentQuantity(ctx context.Context, equipId int) (int, error) {

	res, err := r.db.Query(getEquipmentQuantity, equipId)

	if err != nil {
		log.Println("error while fetching equipment quantity, err : ", err)
		return 0, err
	}

	var quantity int

	res.Next()
	err = res.Scan(&quantity)

	if err != nil {
		log.Println("error while scaning quantity DB row, err : ", err)
		return 0, err
	}

	fmt.Println("quantity: ", quantity)

	return quantity, nil
}

func (r Rental) EquipmentCharges(ctx context.Context, equipId int) (int, error) {

	res, err := r.db.Query(getEquipmentCharges, equipId)

	if err != nil {
		log.Println("error while fetching equipment charges, err : ", err)
		return 0, err
	}

	var rentPerHour int

	res.Next()
	err = res.Scan(&rentPerHour)

	if err != nil {
		log.Println("error while scaning rentcharger DB row, err : ", err)
		return 0, err
	}

	return rentPerHour, nil
}

func (r Rental) CreateBill(ctx context.Context, billing models.Billing) (models.Billing, error) {

	res, err := r.db.Query(createNewBill, billing.Amount, billing.RentId)

	if err != nil {
		log.Println("error while fetching equipment quantity, err : ", err)
		return models.Billing{}, apperrors.ErrDbExce
	}
	res.Next()
	var resp models.Billing

	err = res.Scan(&resp.Id,
		&resp.Date,
		&resp.Amount,
		&resp.RentId)

	if err != nil {
		log.Println("error while scaning billing DB row, err : ", err)
		return models.Billing{}, apperrors.ErrDbParsing
	}

	return resp, nil
}

func (r Rental) UpdateQuantity(ctx context.Context, equipId int, quantity int) error {

	_, err := r.db.Exec(descreaseQuantity, quantity, equipId)

	if err != nil {
		log.Println("error while decrease equipment quantity, err : ", err)
	}

	return nil
}
