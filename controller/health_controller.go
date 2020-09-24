package controller

import (
	"net/http"

	"github.com/andersonlira/market-api/domain"
	"github.com/labstack/echo/v4"
)

//CheckHealth handles the application Health Check
func CheckHealth(c echo.Context) error {
	health := domain.Health{}
	health.Status = "UP"
	return c.JSON(http.StatusOK, health)
}
