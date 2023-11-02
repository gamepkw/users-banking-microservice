package handler

import (
	"net/http"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LoginResponse struct {
	Message      string `json:"message"`
	Token        string `json:"token"`
	IsPinSet     bool   `json:"isPinSet"`
	IsProfileSet bool   `json:"isProfileSet"`
}

func (a *UserHandler) Login(c echo.Context) (err error) {
	time.Sleep(2 * time.Second)
	var user model.User

	if err = c.Bind(&user); err != nil {
		logrus.Errorf("[Login] %s", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if user.Tel == "" || len(user.Tel) != 10 {
		logrus.Errorf("[Login] Invalid Tel")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tel")
	}

	if user.Password == "" {
		logrus.Errorf("[Login] Empty Password")
		return echo.NewHTTPError(http.StatusBadRequest, "Empty Password")
	}

	ctx := c.Request().Context()

	token, isPinSet, isProfileSet, err := a.userService.Login(ctx, &user)
	if err != nil {
		logrus.Errorf("[Login] %s", err)
		return c.JSON(getStatusCode(err), err)
	}

	// token, err := middleware.GenerateJWTToken(user.Tel, 1*time.Hour)
	// if err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusOK, LoginResponse{Message: "Login successful", Token: token, IsPinSet: *isPinSet, IsProfileSet: *isProfileSet})

}
