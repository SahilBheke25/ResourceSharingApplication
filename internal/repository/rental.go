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
	RentEquipment(ctx context.Context, rental models.Rental, availableQantity int, rentPerDay float64) (models.Billing, error)
	EquipmentQuantity(ctx context.Context, equipId int) (int, error)
	EquipmentCharges(ctx context.Context, equipId int) (float64, error)
	// CreateBill(ctx context.Context, billing models.Billing) (models.Billing, error)
	// UpdateQuantity(ctx context.Context, equip_id int, quantity int) error
}

func NewRentalStore(db *sql.DB) RentalStorer {
	return Rental{db: db}
}

func (r Rental) RentEquipment(ctx context.Context, rental models.Rental, availableQantity int, rentPerDay float64) (models.Billing, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("error occured")
		return models.Billing{}, err
	}

	defer func() {
		var txnErr error
		if err != nil {
			txnErr = tx.Rollback()
			if txnErr != nil {
				err = fmt.Errorf("%v %w", err, txnErr)
				return
			}
			log.Println("transaction has been rolled back, err : ", err)
			return
		}

		txnErr = tx.Commit()
		if txnErr != nil {
			err = fmt.Errorf("%v %w", err, txnErr)
			return
		}
	}()

	res := tx.QueryRowContext(ctx, createRental,
		rental.Quantity,
		rental.RentAt,
		rental.RentTill,
		rental.Duration,
		rental.EquipId,
		rental.UserId)

	err = res.Err()
	if err != nil {
		log.Println("error occured while making db reqeust for create reantal, err : ", err)
		return models.Billing{}, err
	}

	var respRental models.Rental
	err = res.Scan(
		&respRental.Id,
		&respRental.Quantity,
		&respRental.RentAt,
		&respRental.RentTill,
		&respRental.Duration,
		&respRental.EquipId,
		&respRental.UserId)
	if err != nil {
		log.Println("error occured scanning created rented details, err : ", err)
		return models.Billing{}, err
	}

	// Update quantity
	err = r.UpdateQuantity(ctx, tx, respRental.EquipId, respRental.Quantity)
	// _, err = r.db.Exec(descreaseQuantity, respRental.Quantity, respRental.EquipId)
	if err != nil {
		log.Println("error while decrease equipment quantity, err : ", err)
		return models.Billing{}, err
	}

	var bill models.Billing
	duration := respRental.Duration
	bill.Amount = duration * rentPerDay
	bill.RentId = respRental.Id

	// create bill
	resp, err := r.CreateBill(ctx, tx, bill)
	if err != nil {
		log.Println("error while calling create bill, err : ", err)
		return models.Billing{}, err
	}

	// Get bill details from transaction-id

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

	return quantity, nil
}

func (r Rental) EquipmentCharges(ctx context.Context, equipId int) (float64, error) {

	res, err := r.db.Query(getEquipmentCharges, equipId)
	if err != nil {
		log.Println("error while fetching equipment charges, err : ", err)
		return 0, err
	}

	var rentPerDay float64

	res.Next()
	err = res.Scan(&rentPerDay)
	if err != nil {
		log.Println("error while scaning rentcharger DB row, err : ", err)
		return 0, err
	}

	return rentPerDay, nil
}

func (r Rental) CreateBill(ctx context.Context, tx *sql.Tx, billing models.Billing) (models.Billing, error) {

	res := tx.QueryRow(createNewBill, billing.Amount, billing.RentId)

	var bill models.Billing

	err := res.Scan(&bill.Id,
		&bill.Date,
		&bill.Amount,
		&bill.RentId)

	if err != nil {
		log.Println("error while scaning billing DB row, err : ", err)
		return models.Billing{}, apperrors.ErrDbParsing
	}

	return bill, nil
}

func (r Rental) UpdateQuantity(ctx context.Context, tx *sql.Tx, equipId int, quantity int) error {

	_, err := tx.Exec(descreaseQuantity, quantity, equipId)
	if err != nil {
		log.Println("error while decrease equipment quantity, err : ", err)
	}

	return nil
}
