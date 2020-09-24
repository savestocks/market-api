package main

import (
	"github.com/andersonlira/market-api/config"
	"github.com/andersonlira/market-api/controller"
	_ "github.com/andersonlira/market-api/gateway/customlog"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
