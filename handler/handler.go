package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/lyoz/sandbox-go/model"
)

// Index ...
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go world.")
}

// Register ...
func Register(c echo.Context) error {
	username := c.FormValue("username")

	err := model.AddUser(username)
	if err != nil {
		c.Logger().Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username")
	}
	return c.String(http.StatusOK, "registered.")
}

// Login ...
func Login(c echo.Context) error {
	username := c.FormValue("username")

	user, err := model.FindUser(username)
	if err != nil {
		c.Logger().Print(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "you are not registered")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

// Hello ...
func Hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+username+"!\n")
}
