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
// @description A little API that shortens URL's
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
	e.GET("/", controller.Root)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/new", controller.NewShortURL)
	e.GET("/:sid", controller.GetFullURL)
	e.Logger.Fatal(e.Start(":8080"))
}
