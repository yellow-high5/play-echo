package handler

import (
	"net/http"
	"play-echo/db"
	"play-echo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		if err := c.Bind(&user); err != nil {
			return err
		}
		if err := c.Validate(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		db := db.GetInstance()
		result := db.Create(&user)
		if result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
		}

		c.Echo().Logger.Debug("create new user✨")
		return c.JSON(http.StatusCreated, user)
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var json map[string]string = map[string]string{}
		if err := c.Bind(&json); err != nil {
			return err
		}
		username := json["username"]
		password := json["password"]

		db := db.GetInstance()
		var user model.User
		db.Where(&model.User{Name: username, Password: password}).Find(&user)
		// Throws unauthorized error
		if user.Name == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		c.Echo().Logger.Debug("get Token🔐")
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

func AuthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.JSON(http.StatusOK, map[string]string{
			"message": name + " is Authenticated✅",
		})
	}
}
