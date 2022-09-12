package controllers

import (
	"net/http"

	"github.com/ali-ghn/go-shop/models"
	r "github.com/ali-ghn/go-shop/repositories"
	"github.com/labstack/echo/v4"
)

type IProductController interface {
	CreateProduct(c echo.Context) error
	ReadProduct(c echo.Context) error
}

type ProductController struct {
	pr r.IProductRepository
}

func NewProductController(pr r.IProductRepository) *ProductController {
	return &ProductController{
		pr: pr,
	}
}

func (pc ProductController) CreateProduct(c echo.Context) error {
	product := models.Product{}
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse the object")
	}
	newProduct, err := pc.pr.CreateProduct(&product)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, please try again")
	}
	return c.JSON(http.StatusCreated, newProduct)
}
