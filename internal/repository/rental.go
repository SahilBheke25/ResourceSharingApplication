package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SahilBheke25/quick-farm-backend/internal/models"
	"github.com/SahilBheke25/quick-farm-backend/internal/pkg/apperrors"
)

const (
	createRental = `INSERT INTO rental(quantity, rent_at, rent_till, duration, equip_id, user_id) 
									VALUES($1, $2, $3, $4, $5, $6)
									RETURNING id, quantity, rent_at, rent_till, duration, equip_id, user_id`

	createNewBill = `INSERT	INTO billing(total_amount, rent_id) VALUES($1, $2)
										RETURNING id, payment_date, total_amount, rent_id`

	getRentedEquipment = ``
)

type Rental struct {
	db *sql.DB
}

type RentalStorer interface {
	RentEquipment(ctx context.Context, rental models.Rental, availableQantity int, rentPerDay float64) (models.Billing, error)
	EquipmentQuantity(ctx context.Context, equipId int) (int, error)
	EquipmentCharges(ctx context.Context, equipId int) (float64, error)
}

func NewRentalStore(db *sql.DB) RentalStorer {
	return Rental{db: db}
}

func (r Rental) RentEquipment(ctx context.Context, rental models.Rental, availableQantity int, rentPerDay float64) (models.Billing, error) {

	// Transaction
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Repo: error occured while starting transaction, err : %v", err)
		return models.Billing{}, apperrors.ErrDbServer
	}

	defer func() {
		var txnErr error
		if err != nil {
			txnErr = tx.Rollback()
			if txnErr != nil {
				err = fmt.Errorf("%v %w", err, txnErr)
				return
			}
			log.Printf("Repo: error transaction has been rolled back, err : %v", err)
			return
		}

		txnErr = tx.Commit()
		if txnErr != nil {
			err = fmt.Errorf("%v %w", err, txnErr)
			log.Printf("Repo: error while commiting transaction, err : %v", err)
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
		log.Printf("Repo: error occured scanning rented details, err : %v", err)
		return models.Billing{}, apperrors.ErrDbServer
	}

	// Update quantity
	err = r.updateQuantity(ctx, tx, respRental.EquipId, respRental.Quantity)
	if err != nil {
		log.Printf("Repo: error while decrease equipment quantity, err : %v", err)
		return models.Billing{}, err
	}

	var bill models.Billing
	duration := respRental.Duration
	bill.Amount = ((duration / 24) * rentPerDay) * float64(respRental.Quantity)
	bill.RentId = respRental.Id

	// create bill
	resp, err := r.createBill(ctx, tx, bill)
	if err != nil {
		log.Printf("Repo: error while calling create bill, err : %v", err)
		return models.Billing{}, err
	}

	return resp, nil
}

func (r Rental) EquipmentQuantity(ctx context.Context, equipId int) (int, error) {

	res, err := r.db.Query(getEquipmentQuantity, equipId)
	if err != nil {
		log.Printf("Repo: error while fetching equipment quantity, err : %v", err)
		return 0, apperrors.ErrDbServer
	}

	var quantity int
	res.Next()
	err = res.Scan(&quantity)
	if err != nil {
		log.Printf("Repo: error while scaning quantity DB row, err : %v", err)
		return 0, apperrors.ErrDbServer
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

func (r Rental) createBill(ctx context.Context, tx *sql.Tx, billing models.Billing) (models.Billing, error) {

	res := tx.QueryRowContext(ctx, createNewBill, billing.Amount, billing.RentId)

	var bill models.Billing
	err := res.Scan(&bill.Id,
		&bill.Date,
		&bill.Amount,
		&bill.RentId)
	if err != nil {
		log.Printf("Repo: error while scaning billing DB row, err : %v", err)
		return models.Billing{}, apperrors.ErrDbServer
	}

	return bill, nil
}

func (r Rental) updateQuantity(ctx context.Context, tx *sql.Tx, equipId int, quantity int) error {

	_, err := tx.Exec(descreaseQuantity, quantity, equipId)
	if err != nil {
		log.Printf("Repo: error while decrease equipment quantity, err : %v", err)
		return apperrors.ErrDbServer
	}

	return nil
}
