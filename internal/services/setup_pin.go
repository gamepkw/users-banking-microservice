package service

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) SetUpPin(c context.Context, u *model.Pin) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	// if err := utils.EncodeBase64(&u.Tel); err != nil {
	// 	return err
	// }

	if err := utils.HashPinBcrypt(&u.Pin); err != nil {
		return err
	}

	if err = a.userRepo.SetUpPin(ctx, u); err != nil {
		return err
	}
	return nil
}
