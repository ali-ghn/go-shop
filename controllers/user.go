package controllers

import (
	"net/http"

	"github.com/ali-ghn/go-shop/repositories"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	ur repositories.IUserRepository
}

func NewUserController(ur repositories.IUserRepository) *UserController {
	return &UserController{
		ur: ur,
	}
}

func Index02(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func (uc UserController) CreateUser(c echo.Context) error {
	return c.String(200, "Hello")
}
