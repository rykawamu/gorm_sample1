package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	//middleware
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
