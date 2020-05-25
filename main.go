package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world.")
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.Start(":3000")
}
