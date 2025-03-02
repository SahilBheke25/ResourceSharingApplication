package user

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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
	UserById(w http.ResponseWriter, r *http.Request)
	EquipmentOwner(w http.ResponseWriter, r *http.Request)
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

	defer r.Body.Close()

	ctx := context.Background()

	// Reading json request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Handler: Error decoding request body:", err)
		utils.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrInvalidReqBody)
		return
	}

	// Validating user data
	err = user.ValidateUser(ctx)
	if err != nil {
		log.Printf("Handler: User validation failed: %v\n", err)
		utils.ErrorResponse(ctx, w, http.StatusBadRequest, err)
		return
	}

	// Create user & checks if user already exist.
	err = u.userService.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Handler: Error calling RegisterUser service: err : %v\n", err)
		switch {
		case errors.Is(err, apperrors.ErrDuplicateUsername):
			utils.ErrorResponse(ctx, w, http.StatusConflict, apperrors.ErrDuplicateUsername)
			return
		case errors.Is(err, apperrors.ErrDuplicateEmail):
			utils.ErrorResponse(ctx, w, http.StatusConflict, apperrors.ErrDuplicateEmail)
			return
		case errors.Is(err, apperrors.ErrDuplicateUid):
			utils.ErrorResponse(ctx, w, http.StatusConflict, apperrors.ErrDuplicateUid)
			return
		default:
			utils.ErrorResponse(ctx, w, http.StatusInternalServerError, apperrors.ErrDbServer)
			return
		}
	}

	utils.SuccessResponse(ctx, w, http.StatusOK, "User Registered Successfully ")
}

func (u *userHandler) UserById(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	// Path param conversion
	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		log.Printf("Handler: error while converting user id param form string to int, err : %v\n", err)
		utils.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrPathParam)
		return
	}

	user, err := u.userService.UserProfile(ctx, userId)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			utils.ErrorResponse(ctx, w, http.StatusNotFound, apperrors.ErrUserNotFound)
			return
		}
		log.Printf("Handler: Unexpected error while calling UserProfile service for user ID %d, err: %v\n", userId, err)
		utils.ErrorResponse(ctx, w, http.StatusInternalServerError, apperrors.ErrInternal)
		return
	}

	utils.SuccessResponse(ctx, w, http.StatusFound, user)
}

func (u *userHandler) EquipmentOwner(w http.ResponseWriter, r *http.Request) {

	// ctx := context.Background()

	// Path param conversion
	// equipId, err := strconv.Atoi(r.PathValue("equip_id"))
	// if err != nil {
	// 	log.Printf("Handler: error while converting user id param form string to int, err : %v\n", err)
	// 	utils.ErrorResponse(ctx, w, http.StatusBadRequest, apperrors.ErrPathParam)
	// 	return
	// }
}
