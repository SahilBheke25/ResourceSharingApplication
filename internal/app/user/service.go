package user

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/pkg/apperrors"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	Authenticate(ctx context.Context, username, password string) (bool, error)
	RegisterUser(ctx context.Context, user models.User) error
	ValidateUser(ctx context.Context, user models.User) (bool, error)
}

func NewService(user repository.UserStorer) Service {

	return service{userRepo: user}

}

func (s service) Authenticate(ctx context.Context, username, password string) (bool, error) {

	resp, err := s.userRepo.GetUserByUsername(ctx, username)

	if err != nil {
		log.Println("error occured while calling getUserByUsername DB opeartion, err : ", err)
		return false, err
	}

	if resp.Password != password {
		return false, apperrors.ErrInvalidCredentials
	}

	return true, nil

}

func (s service) RegisterUser(ctx context.Context, user models.User) error {

	err := s.userRepo.RegisterUser(ctx, user)

	if err != nil {
		log.Println("error occured while calling CreateUser DB opeartion, err : ", err)
		return err
	}
	return nil

}

func (s service) ValidateUser(ctx context.Context, user models.User) (bool, error) {

	if len(user.Username) < 2 {

		return false, fmt.Errorf("username too short")

	} else if len(user.First_name) < 2 {

		return false, fmt.Errorf("first Name too short")

	} else if len(user.Last_name) < 2 {

		return false, fmt.Errorf("last Name too short")

	} else if len(user.Phone) != 10 {

		return false, fmt.Errorf("invalid phone number")

	} else if len(strconv.Itoa(user.Uid)) != 12 {

		return false, fmt.Errorf("invalid UID")

	} else if len(user.Password) < 8 {

		return false, fmt.Errorf("password too short")

	}

	return true, nil
}
