package Repository

import (
	"fmt"
)

const (
	createNewUser  = `INSERT INTO users (user_name, password, first_name, last_name, email, phone, address, pincode, uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	userByusername = `SELECT password from users where user_name = $1`
)

func CreateUser(user_name, password, first_name, last_name, email, phone, address string, pincode int, uid int) error {

	_, err := DB.Exec(createNewUser, user_name, password, first_name, last_name, email, phone, address, pincode, uid)

	if err != nil {
		return fmt.Errorf("Error while creating new user entry: %v", err)
	}

	return nil
}

func AuthenticateUser(userName, password string) (bool, error) {

	var dbPassword string

	err := DB.QueryRow(userByusername, userName).Scan(&dbPassword)

	if err != nil {
		return false, fmt.Errorf("User Not Found: %v", err)
	}

	if dbPassword != password {
		return false, fmt.Errorf("Bad Credentials")
	}

	return true, nil
}
