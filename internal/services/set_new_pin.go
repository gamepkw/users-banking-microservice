package service

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) SetNewPin(c context.Context, u *model.SetNewPin) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	// if err := utils.EncodeBase64(&u.Tel); err != nil {
	// 	return err
	// }

	pinDB, err := a.getHashedPinByUUID(ctx, u.Tel)
	if err != nil {
		return
	}

	if err = utils.ComparePins(pinDB.Pin, u.Pin); err != nil {
		return model.ErrWrongPassword
	}

	if err := utils.HashPinBcrypt(&u.NewPin); err != nil {
		return err
	}

	if err = a.userRepo.SetNewPin(ctx, u); err != nil {
		return err
	}

	return nil
}

func (a *userService) getHashedPinByUUID(c context.Context, uuid string) (res *model.Pin, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if res, err = a.userRepo.GetHashedPinByUUID(ctx, uuid); err != nil {
		return nil, err
	}

	return res, nil
}
