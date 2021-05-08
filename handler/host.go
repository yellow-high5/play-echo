package handler

import (
	"net/http"
	"play-echo/db"
	"play-echo/model"

	"github.com/labstack/echo/v4"
)

func DeleteHost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		db := db.GetInstance()
		var host model.Host
		if err := db.First(&host, id).Error; err != nil {
			return echo.ErrNotFound
		}
		db.Select("Plans").Delete(&host)

		c.Echo().Logger.Debug("delete host successfully")
		return c.JSON(http.StatusNoContent, nil)
	}
}
