package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Root - basic GET request
func Root(c echo.Context) error {
	
	return c.String(http.StatusOK, "Hello URL")
}