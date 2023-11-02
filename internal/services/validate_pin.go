package service

import (
	"context"

	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) ValidatePin(c context.Context, uuid string, pin string) bool {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	// valid := totp.Validate(otp, secretKey)
	pinDB, err := a.getHashedPinByUUID(ctx, uuid)
	if err != nil {
		return false
	}

	if err = utils.ComparePins(pinDB.Pin, pin); err != nil {
		return false
	}

	return true
}
