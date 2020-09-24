package controller

import (
    "net/http"

    "github.com/andersonlira/market-api/gateway/txtdb"
    "github.com/andersonlira/market-api/domain"
	"github.com/labstack/echo/v4"

)


//GetMarketList return all objects 
func GetMarketList(c echo.Context) error {

    list := txtdb.GetMarketList()

	return c.JSON(http.StatusOK, list)
}

func GetMarketByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetMarketByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveMarket(c echo.Context) error {
    it := domain.Market{}
    c.Bind(&it)
    it = txtdb.SaveMarket(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdateMarket(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Market{}
    c.Bind(&it)
    it = txtdb.UpdateMarket(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteMarket(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteMarket(ID)
    return c.JSON(http.StatusOK, result)
}