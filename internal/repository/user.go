package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

const (
	createNewUser = `INSERT INTO users (user_name, password, first_name,
					  	last_name, email, phone, address, pincode, uid)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	userByusername = `SELECT password from users where user_name = $1`
)

type user struct {
	db *sql.DB
}

type UserStorer interface {
	CreateUser(ctx context.Context, user models.User) error
	AuthenticateUser(ctx context.Context, userName, password string) (bool, error)
}

func NewUserStorer(db *sql.DB) UserStorer {
	return user{db: db}
}

func (u user) CreateUser(ctx context.Context, user models.User) error {

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
		return fmt.Errorf("error while creating new user entry: %v", err)
	}

	return nil
}

func (u user) AuthenticateUser(ctx context.Context, userName, password string) (bool, error) {

	var dbPassword string

	err := u.db.QueryRow(userByusername, userName).Scan(&dbPassword)

	if err != nil {
		return false, fmt.Errorf("user Not Found: %v", err)
	}

	if dbPassword != password {
		return false, fmt.Errorf("wrong Password")
	}

	return true, nil
}
