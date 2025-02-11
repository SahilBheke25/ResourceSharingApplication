package user

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/repository"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	AuthenticateUser(ctx context.Context, username, password string) (bool, error)
	CreateUser(ctx context.Context, user models.User) error
	ValidateUser(ctx context.Context, user models.User) (bool, error)
}

func NewService(user repository.UserStorer) Service {

	return service{userRepo: user}

}

func (s service) AuthenticateUser(ctx context.Context, username, password string) (bool, error) {

	resp, err := s.userRepo.AuthenticateUser(ctx, username, password)

	if err != nil {
		log.Println("error occured while calling AuthenticateUser DB opeartion, err : ", err)
		return resp, err
	}
	return resp, nil

}

func (s service) CreateUser(ctx context.Context, user models.User) error {

	err := s.userRepo.CreateUser(ctx, user)

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
