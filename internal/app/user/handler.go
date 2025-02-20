package user

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
)

type userHandler struct {
	userService Service
}

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &userHandler{userService: service}
}

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	// Reading json request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("error while decoding request body, err : ", err)
		utils.ErrorResponse(context.Background(), w, http.StatusBadRequest, apperrors.ErrInvalidReqBody)
		return
	}

	// Authenticating User
	_, err = u.userService.Authenticate(context.Background(), user.Username, user.Password)
	if err != nil {
		log.Println("error while authenticating user credentials, err : ", err)
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			utils.ErrorResponse(context.Background(), w, http.StatusUnauthorized, err)
			return
		} else {
			utils.ErrorResponse(context.Background(), w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.SuccessResponse(context.Background(), w, http.StatusOK, "User Verifed Successfully")
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) {

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
	err = u.userService.RegisterUser(context.Background(), user)

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
