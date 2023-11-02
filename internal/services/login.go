package service

import (
	"context"
	"time"

	"github.com/gamepkw/users-banking-microservice/internal/middleware"
	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
)

func (a *userService) Login(c context.Context, u *model.User) (string, *bool, *bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	userResponse, err := a.getHashedPasswordByUUID(ctx, u.Tel)
	if err != nil {
		return "", nil, nil, err
	}

	if err = utils.ComparePasswords(userResponse.HashedPassword, u.Password); err != nil {
		return "", nil, nil, err
	}

	if err := utils.EncodeBase64(&u.Tel); err != nil {
		return "", nil, nil, err
	}

	token, err := middleware.GenerateJWTToken(u.Tel, 1*time.Hour)
	if err != nil {
		return "", nil, nil, err
	}
	isPinSet := false
	isProfileSet := false
	isPinSet = a.isPinSet(ctx, u.Tel)
	if !isPinSet {
		return token, &isPinSet, &isProfileSet, nil
	}

	isProfileSet = a.isProfileSet(ctx, u.Tel)
	if !isProfileSet {
		return token, &isPinSet, &isProfileSet, nil
	}

	return token, &isPinSet, &isProfileSet, nil
}

func (a *userService) getHashedPasswordByUUID(c context.Context, tel string) (res *model.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if err := utils.EncodeBase64(&tel); err != nil {
		return nil, err
	}

	if res, err = a.userRepo.GetHashedPasswordByUUID(ctx, tel); err != nil {
		return nil, err
	}

	return res, nil
}

func (a *userService) isPinSet(c context.Context, uuid string) bool {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	pinDB, _ := a.getHashedPinByUUID(ctx, uuid)
	if pinDB.Pin == "" {
		return false
	}

	return true
}

func (a *userService) isProfileSet(c context.Context, uuid string) bool {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	isProfileSet, _ := a.userRepo.IsProfileSet(ctx, uuid)

	return isProfileSet
}
