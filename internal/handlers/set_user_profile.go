package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (a *UserHandler) SetUserProfile(c echo.Context) (err error) {
	var user model.UserProfile

	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	user.UUID = c.Get("tel").(string)

	if user.UUID == "" {
		logrus.Errorf("[SetUserProfile] Invalid User")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User")
	}

	ctx := c.Request().Context()

	if err := a.userService.SetUserProfile(ctx, user); err != nil {
		logrus.Errorf("[SetUserProfile] %s", err.Error())
		return c.JSON(getStatusCode(err), ResponseError{Code: "1000", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Message: "Set user profile successful", Body: nil})
}
