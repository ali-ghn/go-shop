package controllers

import (
	"fmt"
	"net/http"

	"github.com/ali-ghn/go-shop/models"
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

func (uc UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	// TODO: Change password storage procedure
	// Use encryption and hash function
	if user.Email == "" || user.Password == "" {
		return c.String(http.StatusBadRequest, "Please make sure you have entered email and password")
	}
	newUser, err := uc.ur.CreateUser(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (uc UserController) ReadUser(c echo.Context) error {
	id := c.QueryParam("id")
	user, err := uc.ur.ReadUser(id)
	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusOK, user)
}

func (uc UserController) UpdateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := uc.ur.UpdateUser(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusOK, newUser)
}
