package handler

import (
	"net/http"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (a *UserHandler) RegisterUser(c echo.Context) (err error) {
	var user model.User

	time.Sleep(5 * time.Second)

	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if user.Tel == "" || len(user.Tel) != 10 {
		logrus.Errorf("[RegisterUser] Invalid Tel")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Tel")
	}

	if user.Password == "" {
		logrus.Errorf("[RegisterUser] Empty Password")
		return echo.NewHTTPError(http.StatusBadRequest, "Empty Password")
	}

	ctx := c.Request().Context()
	res, err := a.userService.RegisterUser(ctx, &user)
	if err != nil {
		logrus.Errorf("[RegisterUser] %s", err.Error())
		return c.JSON(getStatusCode(err), ResponseError{Code: "1000", Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, Response{Message: "Register successful", Body: &res})
}
