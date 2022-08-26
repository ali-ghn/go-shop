package main

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	pc := controllers.NewProductController(getSession())
	e.POST("/product", pc.CreateProduct)
	e.Logger.Fatal(e.Start(":8081"))
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}
