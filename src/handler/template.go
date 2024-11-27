package handler

import (
	"crypto-farm/src/controller"
	pages "crypto-farm/src/pages/home"
	t "crypto-farm/src/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	return pages.Home(pages.HomeProps{}).Render(c.Request().Context(), c.Response())
}

func CreatePot(c echo.Context) error {
	//TODO check how many pots user already has
	//TODO get userID from context
	var userID int64 = 1

	pot := controller.CreatePot(userID)

	return c.JSON(http.StatusOK, pot)
}

func PlantCoin(c echo.Context) error {
	//TODO get userID from context
	var userID int64 = 1
	var plant t.Plant

	if err := c.Bind(&plant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Ticker with this name doesn't exist or date is not correct"})
	}

	price, err := controller.GetPairPrice(plant.Coin, t.USD)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "This ticker doesn't exist"})
	}

	//TODO check that HarvestDate is valid
	plant, err = controller.PlantCoin(
		userID,
		plant.Coin,
		time.Now(),
		plant.HarvestDate,
		price,
	)

	return c.JSON(http.StatusOK, plant)
}
