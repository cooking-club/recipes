package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	setupRoutes(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}

func setupRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/courses/", getCoursesHandler)
	v1.GET("/groups/", getGroupsHandler)
}
