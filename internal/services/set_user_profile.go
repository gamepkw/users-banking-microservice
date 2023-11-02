package service

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (a *userService) SetUserProfile(c context.Context, u model.UserProfile) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err := a.userRepo.SetUserProfile(ctx, u); err != nil {
		return err
	}

	return nil
}
