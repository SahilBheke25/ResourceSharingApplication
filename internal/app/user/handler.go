package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

type userHandler struct {
	userService Service
}

type Handler interface {
	VerifyUserHandler(w http.ResponseWriter, r *http.Request)
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &userHandler{userService: service}
}

func (u *userHandler) VerifyUserHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var user models.User

	// Reading json request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = fmt.Errorf("error while decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Authenticating User
	verified, err := u.userService.AuthenticateUser(context.Background(), user.Username, user.Password)
	if err != nil || !verified {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.HandleResponse(w, "User Verifed Successfully", r)

}

func (u *userHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure body gets closed
	defer r.Body.Close()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
		return
	}

	// Validating user data
	_, err = u.userService.ValidateUser(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	// create user & checks if user already exist.
	err = u.userService.CreateUser(context.Background(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HandleResponse(w, "User Registered Successfully ", r)
}

// ------------------------------------------------------------------

// type user struct {
// 	username string `json: username`
// }

// func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		err = fmt.Errorf("Error while Reading request: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var user user
// 	err = json.Unmarshal(body, &user)

// 	if err != nil {
// 		err = fmt.Errorf("Error while Unmarshelling: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// err = Repository.GetUserByUsername(user.username)

// 	if err != nil {
// 		err = fmt.Errorf("Error while Fetching user by username: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	utils.HandleResponse(w, "Success", r)
// }
