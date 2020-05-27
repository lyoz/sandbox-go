package router

import (
	"net/http"
	"time"

	"github.com/lyoz/sandbox-go/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go world.")
}

func register(c echo.Context) error {
	username := c.FormValue("username")
	err := model.AddUser(username)
	if err != nil {
		c.Logger().Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid username")
	}
	return c.String(http.StatusOK, "registered.\n")
}

func login(c echo.Context) error {
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

func hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+username+"!\n")
}

// Init はEchoのインスタンスの初期設定をするよ
func Init() *echo.Echo {
	e := echo.New()

	e.File("/", "public/index.html")
	e.File("/register", "public/register.html")
	e.POST("/register", register)
	e.File("/login", "public/login.html")
	e.POST("/login", login)

	config := middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(config))
	api.GET("/hello", hello)

	return e
}
