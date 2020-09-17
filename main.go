package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sendMail(to string, subject string, message string) {
	m := gomail.NewMessage()

	m.SetHeader("From", SES.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(SES.Host, SES.Port, SES.Username, SES.Password)
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatal(err)
	}
}

// User ...
type User struct {
	gorm.Model
	Email             string
	Username          string
	PasswordHash      string
	VerifiedAt        sql.NullTime
	VerificationToken string
}

// DB ...
var DB *gorm.DB

// InitDB is initialize db instance
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("echo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&User{})
}

// InitEcho is initialize echo instance
func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.File("/", "public/index.html")
	e.File("/signup", "public/signup.html")
	e.File("/signin", "public/signin.html")

	api := e.Group("/api")
	api.POST("/signup", func(c echo.Context) error {
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")
		var user User
		if err := DB.First(&user, "username=?", username).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%q is already exists", username))
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return echo.ErrInternalServerError
		}
		verificationToken := uuid.New().String()
		DB.Create(&User{Email: email, Username: username, PasswordHash: string(passwordHash), VerificationToken: verificationToken})
		url := "http://localhost:1323/api/verify?token=" + verificationToken
		sendMail(email, "Verify your account", "Follow this link:\n"+url)
		return c.Redirect(http.StatusFound, "/")
	})
	api.POST("/signin", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		var user User
		if err := DB.First(&user, "username=?", username).Error; err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if !user.VerifiedAt.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must verify your account")
		}
		sess, _ := session.Get("session", c)
		sess.Values["signedIn"] = true
		sess.Values["username"] = user.Username
		sess.Save(c.Request(), c.Response())
		return c.String(http.StatusOK, "signed in")
	})
	api.GET("/signout", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["signedIn"] = false
		sess.Save(c.Request(), c.Response())
		return c.String(http.StatusOK, "signed out")
	})
	api.GET("/verify", func(c echo.Context) error {
		token := c.QueryParam("token")
		var user User
		if err := DB.First(&user, "verification_token=?", token).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if user.VerifiedAt.Valid {
			return echo.NewHTTPError(http.StatusBadRequest, "This account is already verified")
		}
		DB.Model(&user).Update("verified_at", time.Now())
		return c.String(http.StatusOK, "verified")
	})

	e.GET("/secret", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		signedIn, ok := sess.Values["signedIn"].(bool)
		if !ok || !signedIn {
			return echo.ErrUnauthorized
		}
		username := sess.Values["username"].(string)
		return c.String(http.StatusOK, "Hello, "+username)
	})

	return e
}

func main() {
	InitDB()

	e := InitEcho()
	e.Logger.Fatal(e.Start(":1323"))
}
