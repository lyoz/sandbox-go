package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index ...
func Index(c echo.Context) error {
	data := "World"
	return c.Render(http.StatusOK, "index", data)
}

// Hello ...
func Hello(c echo.Context) error {
	data := map[string]interface{}{"hello": "world"}
	return c.JSON(http.StatusOK, data)
}
