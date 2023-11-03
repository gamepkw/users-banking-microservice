package service

import (
	"context"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	userRepo "github.com/gamepkw/users-banking-microservice/internal/repositories"
)

type userService struct {
	userRepo       userRepo.UserRepository
	contextTimeout time.Duration
}

func NewUserService(ur userRepo.UserRepository, timeout time.Duration) *userService {
	return &userService{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

type UserService interface {
	RegisterUser(context.Context, *model.User) (model.UserResponse, error)
	Login(c context.Context, u *model.User) (string, *bool, *bool, error)
	SetUpPin(c context.Context, u *model.Pin) (err error)
	SetNewPin(c context.Context, u *model.SetNewPin) (err error)
	SetNewPassword(c context.Context, u *model.UpdatePassword) (res model.UserResponse, err error)
	ResetPassword(c context.Context, u *model.UpdatePassword) (res model.UserResponse, err error)
	ValidatePin(c context.Context, uuid string, pin string) bool
	SetUserProfile(c context.Context, u model.UserProfile) (err error)
}
