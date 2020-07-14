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
	ShortURL string `json:"shortURL" gorm:"primary_key"`
	FullURL  string `json:"fullURL"`
}

// Shorten - returns URLWrapper with shortened url as Slug
func Shorten(url string) (newURL URLWrapper) {
	id, err := helpers.ShortID.Generate()
	if err != nil {
		panic(err)
	}

	newURL.FullURL = url
	newURL.ShortURL = id

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

	ShortenedURL := Shorten(URL.FullURL)

	database.Instance.Create(&ShortenedURL)

	return c.JSON(http.StatusOK, &ShortenedURL)
}

// GetFullURL - get the long url by short
// @Produce  json
// @Success 200 {object} URLWrapper
// @Router /{short_url} [get]
func GetFullURL(c echo.Context) error {
	sid := c.Param("sid")
	URL := URLWrapper{}
	database.Instance.Where(&URLWrapper{ShortURL: sid}).First(&URL)

	return c.JSON(http.StatusOK, &URL)
}
