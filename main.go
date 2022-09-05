package main

import (
	"context"

	"github.com/ali-ghn/go-shop/controllers"
	"github.com/ali-ghn/go-shop/repositories"
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
	uc := controllers.NewUserController(repositories.NewUserRepository(client))
	e.GET("/", controllers.Index)
	e.POST("/user", uc.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
