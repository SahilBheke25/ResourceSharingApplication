package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
)

const (
	createNewUser = `INSERT INTO users 
	(user_name, 
	password, 
	first_name,			
	last_name, 
	email, 
	phone, 
	address, 
	pincode, 
	uid) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	userByusername = `SELECT 
	id, 
	email, 
	pincode, 
	uid, 
	first_name, 
	last_name, 
	user_name, 
	address,
	phone,
	password
	FROM users where user_name = $1`
	userProfile = `SELECT 
	id, 
	email, 
	pincode, 
	uid, 
	first_name, 
	last_name, 
	user_name, 
	address,
	phone 
	FROM users where id = $1`
	equipmentOwner = `SELECT 
	u.id,
  u.user_name, 
  u.first_name, 
  u.last_name, 
  u.phone, 
  u.email, 
  u.address, 
  u.pincode 
	FROM users u
	JOIN equipments e ON u.id = e.user_id
	WHERE e.id = $1`

	updateProfile = `UPDATE users 
	SET user_name = $1, 
	first_name = $2, 
	last_name = $3, 
	email = $4, 
	phone = $5, 
	address = $6, 
	pincode = $7, 
	uid = $8
	WHERE id = $9
	RETURNING id, user_name, first_name, last_name, email, phone, address, pincode, uid;
	`
)

type user struct {
	db *sql.DB
}

type UserStorer interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, userName string) (models.User, error)
	UserProfile(ctx context.Context, userId int) (models.User, error)
	OwnerByEquipmentId(ctx context.Context, equipmentID int) (user models.User, err error)
	UpdateUserProfile(ctx context.Context, updateUser models.User) (models.User, error)
}

func NewUserStorer(db *sql.DB) UserStorer {
	return user{db: db}
}

func (u user) RegisterUser(ctx context.Context, user models.User) error {

	_, err := u.db.Exec(createNewUser,
		user.Username,
		user.Password,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Phone,
		user.Address,
		user.Pincode,
		user.Uid)

	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "unique constraint") || strings.Contains(errMsg, "duplicate key") {
			switch {
			case strings.Contains(errMsg, "username"):
				return apperrors.ErrDuplicateUsername
			case strings.Contains(errMsg, "email"):
				return apperrors.ErrDuplicateEmail
			case strings.Contains(errMsg, "uid"):
				return apperrors.ErrDuplicateUid
			}
		}

		log.Printf("Repo: Failed to create user, err : %v\n", err)
		return apperrors.ErrDbServer
	}

	return nil
}

func (u user) GetUserByUsername(ctx context.Context, userName string) (models.User, error) {

	var user models.User
	err := u.db.QueryRow(userByusername, userName).Scan(&user.Id,
		&user.Email,
		&user.Pincode,
		&user.Uid,
		&user.First_name,
		&user.Last_name,
		&user.Username,
		&user.Address,
		&user.Phone,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error while scanning data, err : %v", err)
			return models.User{}, apperrors.ErrInvalidCredentials
		}
		log.Printf("error while scanning data, err : %v\n", err)
		return models.User{}, apperrors.ErrDbServer
	}

	return user, nil
}

func (u user) UserProfile(ctx context.Context, userId int) (models.User, error) {

	var user models.User
	err := u.db.QueryRow(userProfile, userId).Scan(&user.Id,
		&user.Email,
		&user.Pincode,
		&user.Uid,
		&user.First_name,
		&user.Last_name,
		&user.Username,
		&user.Address,
		&user.Phone,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Repo: User with ID %d not found, err: %v\n", userId, err)
			return models.User{}, apperrors.ErrUserNotFound
		}
		log.Printf("Repo: Database error while fetching user ID %d, err: %v\n", userId, err)
		return models.User{}, apperrors.ErrDbServer
	}

	return user, nil
}

func (u user) OwnerByEquipmentId(ctx context.Context, equipmentID int) (user models.User, err error) {

	err = u.db.QueryRowContext(ctx, equipmentOwner, equipmentID).Scan(
		&user.Id,
		&user.Username,
		&user.First_name,
		&user.Last_name,
		&user.Phone,
		&user.Email,
		&user.Address,
		&user.Pincode,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Repo: No owner found for EquipmentID %d, err: %v\n", equipmentID, err)
			err = apperrors.ErrUserNotFound
			return
		}
		log.Printf("Repo: Database error while fetching equipmentID %d, err: %v\n", equipmentID, err)
		err = apperrors.ErrDbServer
		return
	}

	return
}

func (u user) UpdateUserProfile(ctx context.Context, updateUser models.User) (models.User, error) {

	var updatedUser models.User

	err := u.db.QueryRowContext(ctx, updateProfile,
		updateUser.Username,
		updateUser.First_name,
		updateUser.Last_name,
		updateUser.Email,
		updateUser.Phone,
		updateUser.Address,
		updateUser.Pincode,
		updateUser.Uid,
		updateUser.Id,
	).Scan(
		&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.First_name,
		&updatedUser.Last_name,
		&updatedUser.Email,
		&updatedUser.Phone,
		&updatedUser.Address,
		&updatedUser.Pincode,
		&updatedUser.Uid,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Repo: No user found with ID %d, err: %v\n", updateUser.Id, err)
			return models.User{}, apperrors.ErrUserNotFound
		}
		log.Printf("Repo: Failed to update user ID %d, err: %v\n", updateUser.Id, err)
		return models.User{}, apperrors.ErrDbServer
	}

	return updatedUser, nil
}
