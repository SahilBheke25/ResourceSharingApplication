package user

import (
	"context"
	"errors"
	"log"

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
	UserProfile(ctx context.Context, userId int) (models.User, error)
}

func NewService(user repository.UserStorer) Service {
	return service{userRepo: user}
}

func (s service) Authenticate(ctx context.Context, username, password string) (bool, error) {

	// DB call
	resp, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Service: error occured while calling getUserByUsername DB opeartion, err : %v\n", err)
		return false, err
	}

	// Password Verification
	if resp.Password != password {
		return false, apperrors.ErrInvalidCredentials
	}

	return true, nil
}

func (s service) RegisterUser(ctx context.Context, user models.User) error {

	//DB call
	err := s.userRepo.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Service: error occured while calling CreateUser DB opeartion, err : %v\n", err)
		return err
	}

	return nil
}

func (s service) UserProfile(ctx context.Context, userId int) (user models.User, err error) {

	//DB call
	user, err = s.userRepo.UserProfile(ctx, userId)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return
		}

		log.Printf("Service: Failed to fetch user with ID %d, err: %v\n", userId, err)
		err = apperrors.ErrInternal
		return
	}

	return
}
