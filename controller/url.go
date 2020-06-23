package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// NewShortURL - saves short url to db
func NewShortURL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello URL")
}
