package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) ValidatePin(c echo.Context) (err error) {
	var pin model.Pin

	uuid := c.Get("tel").(string)

	if err = c.Bind(&pin); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	pin.Tel = uuid

	ctx := c.Request().Context()

	if !a.userService.ValidatePin(ctx, pin.Tel, pin.Pin) {
		return c.JSON(http.StatusBadRequest, Response{Message: "Pin is incorrect", Body: nil})
	}

	return c.JSON(http.StatusOK, Response{Message: "Pin is valid", Body: nil})

}
