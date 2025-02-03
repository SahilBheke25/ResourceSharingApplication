package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/handle"
	Models "github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	Repository "github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

// const (
// 	EncryptionKey = "thisis32bitlongpassphrase!"
// )

func Register(w http.ResponseWriter, r *http.Request) {
	// Ensure body gets closed
	defer r.Body.Close()

	var user Models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
		return
	}

	_, err = validateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// create user & checks if user already exist.
	err = Repository.CreateUser(user.Username, user.Password, user.First_name, user.Last_name, user.Email, user.Phone, user.Address, user.Pincode, user.Uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handle.HandleResponse(w, "User Registered Successfully ", r)
}

func validateUser(user Models.User) (bool, error) {
	if len(user.Username) < 2 {

		return false, fmt.Errorf("Username too short")

	} else if len(user.First_name) < 2 {

		return false, fmt.Errorf("First Name too short")

	} else if len(user.Last_name) < 2 {

		return false, fmt.Errorf("Last Name too short")

	} else if len(user.Phone) != 10 {

		return false, fmt.Errorf("Invalid phone number")

	} else if len(strconv.Itoa(user.Uid)) != 12 {

		return false, fmt.Errorf("Invalid UID")

	} else if len(user.Password) < 8 {

		return false, fmt.Errorf("Password too short")

	}

	return true, nil
}
