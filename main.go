package main

import (
	"context"

	"github.com/ali-ghn/go-shop/controllers"
	"github.com/ali-ghn/go-shop/repositories"
	"github.com/ali-ghn/go-shop/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
}

func main() {
	e := echo.New()
	uc := controllers.NewUserController(repositories.NewUserRepository(client), services.Auth{
		AuthKey: []byte("Some Random Key"),
	})
	pc := controllers.NewProductController(repositories.NewProductRepository(client))
	e.GET("/", controllers.Index)
	e.POST("/user", uc.CreateUser)
	e.PUT("/user", uc.UpdateUser)
	e.GET("/user", uc.ReadUser)
	e.POST("/SignIn", uc.SignIn)
	e.POST("/product", pc.CreateProduct)
	e.Logger.Fatal(e.Start(":8081"))
}
