package router

import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/lyoz/sandbox-go/constant"
	"github.com/lyoz/sandbox-go/handler"
)

// Init はEchoのインスタンスの初期設定をするよ
func Init() *echo.Echo {
	e := echo.New()

	e.File("/", "public/index.html")
	e.File("/register", "public/register.html")
	e.POST("/register", handler.Register)
	e.File("/login", "public/login.html")
	e.POST("/login", handler.Login)

	e.File("test", "public/test.html")

	config := middleware.JWTConfig{
		SigningKey: []byte(constant.SigningKey),
	}

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(config))
	api.GET("/hello", handler.Hello)

	return e
}
