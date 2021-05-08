package main

import (
	"net/http"
	"play-echo/db"
	"play-echo/handler"
	customMiddleware "play-echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func health(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func main() {
	// migrate db
	db.Migrate()

	// set middleware
	e := echo.New()
	e.Debug = true
	e.Validator = customMiddleware.NewValidator()
	e.Use(customMiddleware.SetLogger())
	e.Use(customMiddleware.SetBodyDump(e.Logger))

	// routing
	e.GET("/health", health)
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handler.AuthCheck())
	r.GET("/users/:id", handler.GetUser())
	r.GET("/users", handler.GetAllUser())
	r.PUT("/users/:id", handler.UpdateUser())
	r.DELETE("/users", handler.DeleteUser())

	e.GET("/plans", handler.GetPlan())
	e.DELETE("/hosts/:id", handler.DeleteHost())

	// boost echo server
	e.Logger.SetHeader(customMiddleware.DebugLogFormat())
	e.Logger.Fatal(e.Start(":1323"))
}
