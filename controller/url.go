package controller

import (
	"net/http"

	"github.com/eidyz/short-url/database"
	"github.com/eidyz/short-url/helpers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// URLWrapper ---
type URLWrapper struct {
	gorm.Model
	Short string `json:"short" gorm:"primary_key"`
	Long  string `json:"long"`
}

// Shorten - returns URLWrapper with shortened url as Slug
func Shorten(url string) (newURL URLWrapper) {
	id, err := helpers.ShortID.Generate()
	if err != nil {
		panic(err)
	}

	newURL.Long = url
	newURL.Short = id

	return newURL
}

// NewShortURL - saves short url to db
// @Summary Shortens URL
// @Accept  json
// @Produce  json
// @Success 200 {object} URLWrapper
// @Router /new [post]
func NewShortURL(c echo.Context) error {
	URL := URLWrapper{}
	if err := c.Bind(&URL); err != nil {
		return err
	}

	ShortenedURL := Shorten(URL.Long)

	database.Instance.Create(&ShortenedURL)

	return c.JSON(http.StatusOK, &ShortenedURL)
}

// GetLongURL - get the long url by short
// @Produce  json
// @Success 200 {object} URLWrapper
// @Router /{short_url} [get]
func GetLongURL(c echo.Context) error {
	short := c.Param("short")
	URL := URLWrapper{}
	database.Instance.Where(&URLWrapper{Short: short}).First(&URL)

	return c.JSON(http.StatusOK, &URL)
}
