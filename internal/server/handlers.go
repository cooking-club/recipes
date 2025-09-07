package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/cooking-club/recipes/internal/schedule"
)

type Day struct {
	Date    string            `json:"date"`
	Records []schedule.Record `json:"records"`
}

func getCoursesHandler(c echo.Context) error {
	f := c.QueryParam("f")
	y := c.QueryParam("y")
	g := c.QueryParam("g")
	s := c.QueryParam("s")
	w := c.QueryParam("w")

	perp := true
	n, _ := strconv.Atoi(w)
	week := "0"

	if n%4 == 0 || n%4 == 2 {
		week = "1"
	}

	if n < 2 {
		perp = false
	}

	_, data := schedule.GetSchedule(f, y, g, s, week, perp)
	return c.JSON(http.StatusOK, data)
}
