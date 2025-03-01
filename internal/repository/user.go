package repository

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
)

const (
	createNewUser = `INSERT INTO users (user_name, password, first_name,
					  			last_name, email, phone, address, pincode, uid)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	userByusername = `SELECT password from users where user_name = $1`
	userProfile    = `SELECT user_name, first_name, last_name from users where id = $1`
)

type user struct {
	db *sql.DB
}

type UserStorer interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, userName string) (models.UserCredentials, error)
	// UserProfile(ctx context.Context, userId int) (models.User, error)
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
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "username") {
				return apperrors.ErrDuplicateUsername
			}
			if strings.Contains(err.Error(), "email") {
				return apperrors.ErrDuplicateEmail
			}
		}
		log.Printf("Failed to create user: %v", err)
		return apperrors.ErrDbServer
	}

	return nil
}

func (u user) GetUserByUsername(ctx context.Context, userName string) (models.UserCredentials, error) {

	var user models.UserCredentials

	err := u.db.QueryRow(userByusername, userName).Scan(&user.Password)
	if err == sql.ErrNoRows {
		log.Println("error while scanning data, err : ", err)
		return models.UserCredentials{}, apperrors.ErrInvalidCredentials
	}
	if err != nil {
		log.Println("error while scanning data, err : ", err)
		return models.UserCredentials{}, apperrors.ErrDbServer
	}

	return user, nil
}

// func (u user) UserProfile(ctx context.Context, userId int) (models.User, error) {

// 	var user models.UserProfile
// 	err := u.db.QueryRow(userProfile, userId).Scan(&user.Username, &user.First_name, &user.Last_name)

// }
