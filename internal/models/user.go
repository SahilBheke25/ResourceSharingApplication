package models

import (
	"context"
	"fmt"
	"strconv"
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Pincode    int    `json:"pincode"`
	Uid        int    `json:"uid"`
}

type UserProfile struct {
	Username   string `json:"username"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
}

type UserCredentials struct {
	Password string `json:"password"`
}

func (u User) ValidateUser(ctx context.Context, user User) (bool, error) {

	if len(user.Username) < 2 {
		return false, fmt.Errorf("username too short")
	}
	if len(user.First_name) < 2 {
		return false, fmt.Errorf("first Name too short")
	}
	if len(user.Last_name) < 2 {
		return false, fmt.Errorf("last Name too short")
	}
	if len(user.Phone) != 10 {
		return false, fmt.Errorf("invalid phone number")
	}
	if len(strconv.Itoa(user.Uid)) != 12 {
		return false, fmt.Errorf("invalid UID")
	}
	if len(user.Password) < 8 {
		return false, fmt.Errorf("password too short")
	}
	return true, nil
}
