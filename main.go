package main

import (
	"fmt"

	"github.com/eidyz/short-url/controller"
	"github.com/eidyz/short-url/database"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected to Database!")

	e := echo.New()
	e.POST("/", controller.NewShortURL)
	e.Logger.Fatal(e.Start(":8080"))
}
