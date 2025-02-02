package Repository

import (
	"fmt"
)

const (
	createNewUser = `INSERT INTO users (user_name, password, first_name, last_name, email, phone, address, pincode, uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	userByid      = `SELECT * from users where user_name = $1`
)

func CreateUser(user_name, password, first_name, last_name, email, phone, address string, pincode int, uid int) error {

	_, err := DB.Exec(createNewUser, user_name, password, first_name, last_name, email, phone, address, pincode, uid)

	if err != nil {
		return fmt.Errorf("Error while creating new user entry: %v", err)
	}

	return nil
}

func GetUserByUsername(userName string) error {
	res, err := DB.Exec(userByid, userName)

	if err != nil {
		return fmt.Errorf("Error while feting User by id: %v", err)
	}

	fmt.Printf("%v \n", res)

	return nil
}
