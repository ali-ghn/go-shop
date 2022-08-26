package controllers

import (
	"log"
	"myapp/models"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
)

type ProductController struct {
	productRepository *repositories.ProductRepository
}

func NewProductController(session *mgo.Session) *ProductController {
	return &ProductController{
		productRepository: repositories.NewProductRepository(session),
	}
}

func (pc ProductController) CreateProduct(c echo.Context) error {
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	log.Println(p)
	err := pc.productRepository.CreateProduct(p.Name, p.Price)
	return err
}
