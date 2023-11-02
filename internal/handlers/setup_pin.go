package handler

import (
	"net/http"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
	"github.com/gamepkw/users-banking-microservice/internal/utils"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) SetUpPin(c echo.Context) (err error) {
	var pin model.Pin

	expectedTel := c.Get("tel").(string)

	if err = c.Bind(&pin); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := utils.EncodeBase64(&pin.Tel); err != nil {
		return err
	}

	if pin.Tel != expectedTel {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Tel mismatch"})
	}

	if pin.Pin == "" || len(pin.Pin) != 6 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Pin")
	}

	ctx := c.Request().Context()

	if err = a.userService.SetUpPin(ctx, &pin); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Message: "Set pin successfully", Body: nil})
}
