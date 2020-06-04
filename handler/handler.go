package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IndexGet ...
func IndexGet(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

// HelloGet ...
func HelloGet(c echo.Context) error {
	greetingto := c.QueryParam("greetingto")
	return c.Render(http.StatusOK, "hello", greetingto)
}

// HelloPost ...
func HelloPost(c echo.Context) error {
	greetingto := c.FormValue("greetingto")
	return c.Render(http.StatusOK, "hello", greetingto)
}

// HelloFormGet ...
func HelloFormGet(c echo.Context) error {
	return c.Render(http.StatusOK, "hello_form", nil)
}

// APIHelloGet ...
func APIHelloGet(c echo.Context) error {
	data := map[string]interface{}{"hello": "world"}
	return c.JSON(http.StatusOK, data)
}

// APIHelloPost ...
func APIHelloPost(c echo.Context) error {
	param := map[string]interface{}{}
	c.Bind(&param)
	data := map[string]interface{}{"hello": param["greetingto"]}
	return c.JSON(http.StatusOK, data)
}
