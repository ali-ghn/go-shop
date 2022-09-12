package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ali-ghn/go-shop/models"
	"github.com/ali-ghn/go-shop/repositories"
	"github.com/ali-ghn/go-shop/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	ur   repositories.IUserRepository
	auth services.IAuth
}

func NewUserController(ur repositories.IUserRepository, auth services.IAuth) *UserController {
	return &UserController{
		ur:   ur,
		auth: auth,
	}
}

func (uc UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	// TODO: Change password storage procedure
	// Use encryption and hash function
	// TODO: Add email checking for duplicate emails
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
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.String(http.StatusForbidden, "You don't have access to this endpoint")
	}
	claims, err := uc.auth.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Error while authorizing")
	}

	currentUser, err := uc.ur.ReadUserByEmail(claims.Email)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}

	var isAuthorized bool

	for _, v := range currentUser.Roles {
		if v == "Admin" {
			isAuthorized = true
		}
	}
	if !isAuthorized {
		return c.String(http.StatusForbidden, "You don't have access to this endpoint")

	}
	id := c.QueryParam("id")
	user, err := uc.ur.ReadUser(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusOK, user)
}

func (uc UserController) UpdateUser(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.String(http.StatusForbidden, "You don't have access to this endpoint")
	}
	claims, err := uc.auth.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Error while authorizing")
	}

	currentUser, err := uc.ur.ReadUserByEmail(claims.Email)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}

	var isAuthorized bool

	for _, v := range currentUser.Roles {
		if v == "Admin" {
			isAuthorized = true
		}
	}
	if !isAuthorized {
		return c.String(http.StatusForbidden, "You don't have access to this endpoint")

	}

	user := models.User{}
	c.Bind(&user)
	newUser, err := uc.ur.UpdateUser(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusOK, newUser)
}

func (uc UserController) SignIn(c echo.Context) error {
	userCredential := models.SignInRequest{}
	c.Bind(&userCredential)
	user, err := uc.ur.ReadUserByEmail(userCredential.Email)
	if err != nil {
		fmt.Println("Something went wrong: %w", err)
		c.String(http.StatusInternalServerError, "Couldn't sign in")
		return err
	}
	if userCredential.Password != user.Password {
		c.String(http.StatusBadRequest, "Couldn't sign in")
		return fmt.Errorf("couldn't sign in")
	}
	userClaims := services.UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		Email: userCredential.Email,
	}
	st, err := uc.auth.CreateToken(&userClaims)
	if err != nil {
		c.String(http.StatusInternalServerError, "something went wrong")
		return err
	}
	c.String(http.StatusOK, st)
	return nil
}
