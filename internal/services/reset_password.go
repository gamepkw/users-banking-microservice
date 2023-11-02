package service

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) ResetPassword(c context.Context, u *model.UpdatePassword) (res model.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err = utils.EncodeBase64(&u.Tel); err != nil {
		return
	}

	if err = utils.HashPasswordBcrypt(&u.Password); err != nil {
		return
	}

	if utils.ValidatePassword(u.NewPassword) {

		if err := utils.HashPasswordBcrypt(&u.NewPassword); err != nil {
			return res, err
		}
		res, err = a.userRepo.SetPassword(ctx, u)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
