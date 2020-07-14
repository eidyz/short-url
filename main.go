package main

import (
	"fmt"

	"github.com/eidyz/short-url/controller"
	"github.com/eidyz/short-url/database"
	_ "github.com/eidyz/short-url/docs"
	"github.com/eidyz/short-url/helpers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Short URL API
// @version 1.0
// @description Shortens given URL
// @termsOfService http://swagger.io/terms/

// @BasePath /v2
func main() {
	helpers.InitShortID()
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected to Database!")

	database.Instance.AutoMigrate(&controller.URLWrapper{})

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", controller.Root)
	e.POST("/new", controller.NewShortURL)
	e.GET("/:short", controller.GetLongURL)
	e.Logger.Fatal(e.Start(":8080"))
}
