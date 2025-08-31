package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
