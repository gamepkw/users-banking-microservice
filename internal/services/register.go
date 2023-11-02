package service

import (
	"context"

	"github.com/gamepkw/users-banking-microservice/internal/utils"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (a *userService) RegisterUser(c context.Context, u *model.User) (res model.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	u.UUID = u.Tel

	if err := utils.EncodeBase64(&u.UUID); err != nil {
		return res, err
	}

	if utils.ValidatePassword(u.Password) {

		if err := utils.HashPasswordBcrypt(&u.Password); err != nil {
			return res, err
		}
		res, err = a.userRepo.RegisterUser(ctx, u)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

