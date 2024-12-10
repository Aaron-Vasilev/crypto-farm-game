package handler

import (
	"crypto-farm/src/auth"
	"crypto-farm/src/components"
	"crypto-farm/src/consts"
	"crypto-farm/src/controller"
	pages "crypto-farm/src/pages/home"
	t "crypto-farm/src/types"
	"crypto-farm/src/utils"
	"fmt"
	"net/http"

	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	userId := auth.GetUserIDFromCtx(c)

	plants, err := controller.GetPotsWithPlants(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Something is wrong"})
	}

	return pages.Home(plants).Render(c.Request().Context(), c.Response())
}

func CreatePot(c echo.Context) error {
	//TODO check how many pots user already has
	//TODO get userID from context
	var userID int64 = 1

	pot := controller.CreatePot(userID)

	return c.JSON(http.StatusOK, pot)
}

type plantCoin struct {
	Coin    t.Ticker
	PlantID int
}

func PlantCoin(c echo.Context) error {
	userId := auth.GetUserIDFromCtx(c)
	coin := t.Ticker(c.Request().FormValue("coin"))
	potId, err := strconv.Atoi(c.Request().FormValue("potId"))

	//TODO check that pot is empty
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No potId in the request"})
	}

	price, err := controller.GetPairPrice(coin, t.USD)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "This ticker doesn't exist"})
	}

	plant := consts.Plants[coin]
	plantTime := time.Now().UTC()
	harvestTime := plantTime.Add(time.Minute * time.Duration(plant.Minutes))

	pot, err := controller.PlantCoin(
		userId,
		potId,
		coin,
		harvestTime,
		harvestTime,
		price,
	)

	return components.Pot(t.PotWithPlant{
		Pot:   pot,
		Plant: plant,
	}).Render(c.Request().Context(), c.Response())
}

func CheckPlant(c echo.Context) error {
	userId := auth.GetUserIDFromCtx(c)
	potId, err := strconv.Atoi(c.Param("potId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No potId in the request"})
	}

	pot, err := controller.GetPotById(userId, potId)

	if err != nil {
		fmt.Println("✡️  line 92 err", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The pot doesn't exist"})
	}

	if utils.PlantIsReady(pot) {
		price, err := controller.GetPairPrice(pot.Coin, t.USD)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "This ticker doesn't exist"})
		} else {
			controller.UpdateHarvestPrice(userId, potId, price)

			pot.PlantPrice = price
		}
	}

	return components.PotModal(t.PotWithPlant{
		Pot:   pot,
		Plant: consts.Plants[pot.Coin],
	}).Render(c.Request().Context(), c.Response())
}
