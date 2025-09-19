package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cooking-club/recipes/internal/groups"
	"github.com/cooking-club/recipes/internal/schedule"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getCoursesHandler(c echo.Context) error {
	g := c.QueryParam("g")
	d := c.QueryParam("d")

	if g == "" || d == "" {
		return c.JSON(http.StatusBadRequest, Error{1, "no params"})
	}

	group, err := strconv.Atoi(g) // group id
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{1, "bad g"})
	}

	date, err := strconv.ParseInt(d, 10, 64) // unix seconds
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{1, "bad d"})
	}

	start, _ := time.Parse(time.RFC3339, "2025-09-01T00:00:00Z")
	week := time.Unix(date, 0).Sub(start) / (time.Hour * 24 * 7)

	var p, l = 0, 0
	if week%2 == 0 {
		p = 7 * 6
		l = 7 * 6 * 2
	} else {
		p = 0
		l = 7 * 6
	}

	data, err := schedule.GetSchedule(group, p, l)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func getGroupsHandler(c echo.Context) error {
	data, err := groups.GetGroups()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{1, "server is mid"})
	}

	return c.JSON(http.StatusOK, data)
}
