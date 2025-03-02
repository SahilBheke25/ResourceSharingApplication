package models

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Pincode    int    `json:"pincode"`
	Uid        int    `json:"uid,omitempty"`
}

// type UserProfile struct {
// 	Username   string `json:"username"`
// 	First_name string `json:"firstname"`
// 	Last_name  string `json:"lastname"`
// }

type UserCredentials struct {
	Password string `json:"password"`
}

func (u User) ValidateUser(ctx context.Context, validatePassword bool) error {

	var validationErrors []string

	// Regex patterns
	phoneRegex := regexp.MustCompile(`^\d{10}$`)
	uidRegex := regexp.MustCompile(`^\d{12}$`)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Validations
	if len(u.Username) < 2 {
		validationErrors = append(validationErrors, "username must be at least 2 characters")
	}
	if len(u.First_name) < 2 {
		validationErrors = append(validationErrors, "first name must be at least 2 characters")
	}
	if len(u.Last_name) < 2 {
		validationErrors = append(validationErrors, "last name must be at least 2 characters")
	}
	if !phoneRegex.MatchString(u.Phone) {
		validationErrors = append(validationErrors, "phone number must be exactly 10 digits")
	}
	if !uidRegex.MatchString(fmt.Sprintf("%012d", u.Uid)) {
		validationErrors = append(validationErrors, "UID must be exactly 12 digits")
	}
	if u.Email == "" || !emailRegex.MatchString(u.Email) {
		validationErrors = append(validationErrors, "invalid email format")
	}
	if validatePassword && len(u.Password) < 8 {
		validationErrors = append(validationErrors, "password must be at least 8 characters")
	}
	if !(u.Pincode >= 100000 && u.Pincode <= 999999) {
		validationErrors = append(validationErrors, "pincode must be exactly 6 digits")
	}

	// Return all errors
	if len(validationErrors) > 0 {
		return fmt.Errorf(strings.Join(validationErrors, ", "))
	}
	return nil
}
