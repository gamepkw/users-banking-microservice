package service

import (
	"context"
	"fmt"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) SetNewPassword(c context.Context, u *model.UpdatePassword) (res model.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	userResponse, err := a.getHashedPasswordByUUID(ctx, u.Tel)
	if err != nil {
		return res, err
	}

	if err = utils.ComparePasswords(userResponse.HashedPassword, u.Password); err != nil {
		return res, err
	}

	if err = utils.EncodeBase64(&u.Tel); err != nil {
		return
	}

	if err = utils.HashPasswordBcrypt(&u.NewPassword); err != nil {
		return
	}

	fmt.Println(u.NewPassword)

	res, err = a.userRepo.SetPassword(ctx, u)
	if err != nil {
		return res, err
	}

	return res, nil
}
