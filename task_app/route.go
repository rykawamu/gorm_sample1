package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"taskapp/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	//middleware
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Static
	e.Static("/assets", "public/assets")

	//Routing
	e.File("/", "public/index.html")
	e.File("/tasks", "public/tasks.html")

	/*
		e.GET("/login", func(c echo.Context) error {
			return c.String(http.StatusOK, "login page")
		})
	*/
	e.File("/login", "public/login.html")
	e.POST("/login", handler.Login)

	e.GET("/signup", func(c echo.Context) error {
		return c.String(http.StatusOK, "signup page")
	})

	/*
		e.GET("/tasks", func(c echo.Context) error {
			return c.String(http.StatusOK, "tasks page")
		})
	*/

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/tasks", handler.GetTasks)

	return e
}
