package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) VerifyUser(c echo.Context) (err error) {
	var set model.UpdatePassword

	if err = c.Bind(&set); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	if !a.AuthService.ValidateOtp(ctx, set.Tel, set.Otp) {
		return c.JSON(http.StatusBadRequest, Response{Message: "Otp is invalid", Body: nil})
	}

	return c.JSON(http.StatusOK, Response{Message: "Otp is valid", Body: nil})

}
