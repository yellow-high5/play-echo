package handler

import (
	"net/http"
	"play-echo/db"
	"play-echo/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/imdario/mergo"
	"github.com/labstack/echo/v4"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		db := db.GetInstance()
		var user model.User
		db.First(&user, id)

		if user.Name == "" {
			return echo.ErrNotFound
		}

		c.Echo().Logger.Debug("get user successfully")
		return c.JSON(http.StatusOK, user)
	}
}

func GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))

		var users []model.User
		db := db.GetInstance()
		db.Limit(limit).Find(&users)

		c.Echo().Logger.Debug("get user list successfully")
		return c.JSON(http.StatusOK, users)
	}
}

func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateUser model.User
		if err := c.Bind(&updateUser); err != nil {
			return err
		}

		id := c.Param("id")
		db := db.GetInstance()
		var user model.User
		if err := db.First(&user, id).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "user not find")
		}

		mergo.Merge(&updateUser, user)

		if err := c.Validate(&updateUser); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		db.Where(user).Updates(updateUser)

		c.Echo().Logger.Debug("udpate user successfully")
		return c.JSON(http.StatusNoContent, nil)
	}
}

func DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["name"].(string)
		db := db.GetInstance()
		var user model.User
		if err := db.First(&user, "name = ?", username).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "user not find")
		}
		db.Delete(&user)

		c.Echo().Logger.Debug("delete user successfully")
		return c.JSON(http.StatusNoContent, nil)
	}
}
