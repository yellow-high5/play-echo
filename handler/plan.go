package handler

import (
	"net/http"
	"play-echo/db"
	"play-echo/model"

	"github.com/labstack/echo/v4"
)

func GetPlan() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := db.GetInstance()
		var host model.Host

		if err := db.Preload("Plans").First(&host).Error; err != nil {
			return echo.ErrNotFound
		}

		c.Echo().Logger.Debug("get plan successfully")
		return c.JSON(http.StatusOK, host)
	}
}
