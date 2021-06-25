package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MainCtrl struct {
}

func (m *MainCtrl) HealthCheck(c echo.Context) error {
	resp := map[string]interface{}{}
	resp["status"] = "ok"
	resp["service"] = "Music chart API"

	return c.JSON(http.StatusOK, resp)
}
