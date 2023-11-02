package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	authModel "github.com/gamepkw/authentication-banking-microservice/models"
	model "github.com/gamepkw/users-banking-microservice/internal/models"
	userService "github.com/gamepkw/users-banking-microservice/internal/services"

	"github.com/gamepkw/users-banking-microservice/internal/middleware"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService userService.UserService
	AuthService authModel.AuthenticationService
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Message string              `json:"message"`
	Body    *model.UserResponse `json:"body,omitempty"`
}

func NewUserHandler(e *echo.Echo, us userService.UserService, auths authModel.AuthenticationService) {
	handler := &UserHandler{
		userService: us,
		AuthService: auths,
	}

	restrictedGroup := e.Group("/users/pin")
	restrictedGroup.Use(middleware.CustomJWTMiddleware)

	e.POST("/users/register", handler.RegisterUser)
	e.POST("/users/login", handler.Login)
	e.POST("/users/set-new-password", handler.ResetPassword)
	restrictedGroup.PUT("/set-pin", handler.SetUpPin)
	restrictedGroup.PUT("/set-new-pin", handler.SetNewPin)
	restrictedGroup.POST("/verify-pin", handler.ValidatePin)
}

var TimestampFormat = "2006-01-02 15:04:05"

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
