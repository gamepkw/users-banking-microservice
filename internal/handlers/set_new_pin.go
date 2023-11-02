package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) SetNewPin(c echo.Context) (err error) {
	var pin model.SetNewPin

	expectedTel := c.Get("tel").(string)

	if err = c.Bind(&pin); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := utils.EncodeBase64(&pin.Tel); err != nil {
		return err
	}

	if pin.Tel != expectedTel {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Unauthorized"})
	}

	if pin.Pin == "" || len(pin.Pin) != 6 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Pin")
	}

	if pin.NewPin == "" || len(pin.NewPin) != 6 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid New Pin")
	}
	if pin.NewPin == pin.Pin {
		return echo.NewHTTPError(http.StatusBadRequest, "New pin same as old pin")
	}

	ctx := c.Request().Context()

	if err = a.userService.SetNewPin(ctx, &pin); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: "Set new pin failed"})
	}

	return c.JSON(http.StatusOK, Response{Message: "Set new pin successfully", Body: nil})
}
