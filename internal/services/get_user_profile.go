package service

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) GetUserProfile(c context.Context, u *model.User) (res model.UserResponse, err error) {
	_, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err := utils.EncodeBase64(&u.Tel); err != nil {
		return res, err
	}

	// if res, err = a.userRepo.GetUserProfile(ctx, tel); err != nil {
	// 	return nil, err
	// }

	return res, nil
}
