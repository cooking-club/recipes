package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/get/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
}
