package handler

import (
	"crypto-farm/src/auth"
	"crypto-farm/src/controller"
	pages "crypto-farm/src/pages/home"
	t "crypto-farm/src/types"
	"fmt"
	"net/http"
	"strconv"
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
	//TODO check that HarvestDate is valid

	price, err := controller.GetPairPrice(plant.Coin, t.USD)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "This ticker doesn't exist"})
	}

	plant, err = controller.CreateCoin(
		userID,
		plant.Coin,
		time.Now(),
		plant.HarvestDate,
		price,
	)

	return c.JSON(http.StatusOK, plant)
}

func isPlantReady(plant t.Plant) bool {
	// TODO
	return true
}

func HarvestCoin(c echo.Context) error {
	userId := auth.GetUserIDFromCtx(c)
	plantId, err := strconv.Atoi(c.QueryParam("plantId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No plantId in the request"})
	}

	plant, err := controller.GetPlant(userId, plantId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "The plant doesn't exist",
			"userId":  fmt.Sprintf("%d", userId),
			"plantId": c.QueryParam("plantId"),
		})
	}

	if isPlantReady(plant) {
		harvestPrice, err := controller.GetPairPrice(plant.Coin, t.USD)

		if err != nil {
			fmt.Println("✡️  line 84 err", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Oops, Something went wrong"})
		}
		profit := harvestPrice - plant.PlantPrice

		plant, err = controller.HarvestPlant(userId, plantId, harvestPrice, profit)

		if err != nil {
			fmt.Println("✡️  line 92 err", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Oops, Something went wrong"})
		} else {
			return c.JSON(http.StatusOK, plant)
		}
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Plant is not ready"})
	}
}
