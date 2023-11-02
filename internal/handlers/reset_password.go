package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (a *UserHandler) ResetPassword(c echo.Context) (err error) {
	var user model.UpdatePassword

	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if user.Tel == "" || len(user.Tel) != 10 {
		logrus.Errorf("[ResetPassword] Invalid Tel")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tel")
	}

	if user.NewPassword == "" {
		logrus.Errorf("[ResetPassword] Empty Password")
		return echo.NewHTTPError(http.StatusBadRequest, "Empty Password")
	}

	if user.Password == user.NewPassword {
		logrus.Errorf("[ResetPassword] New password can not be the cuurent password")
		return echo.NewHTTPError(http.StatusBadRequest, "New password can not be the cuurent password")
	}

	ctx := c.Request().Context()

	res, err := a.userService.ResetPassword(ctx, &user)
	if err != nil {
		logrus.Errorf("[ResetPassword] %s", err)
		return c.JSON(getStatusCode(err), err)
	}

	return c.JSON(http.StatusOK, Response{Message: "Set new password successful", Body: &res})

}
