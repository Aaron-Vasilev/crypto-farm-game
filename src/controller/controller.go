package controller

import (
	"crypto-farm/src/db"
	t "crypto-farm/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreatePot(userId int64) t.Pot {
	pot := t.Pot{
		UserID: userId,
	}
	query := "INSERT INTO farm.pot (user_id) VALUES ($1) RETURNING id;"

	err := db.DB.QueryRow(query, userId).Scan(&pot.ID)

	if err != nil {
		log.Fatalf("Failed to create pot: %v", err)
	}

	return pot
}

func CreateCoin(
	userId int64,
	coin t.Ticker,
	plantDate, harvestDate time.Time,
	plantPrice float32,
) (t.Plant, error) {
	plant := t.Plant{
		UserID:      userId,
		Coin:        coin,
		PlantDate:   plantDate,
		HarvestDate: harvestDate,
		PlantPrice:  plantPrice,
	}
	query := "INSERT INTO farm.plant (user_id, coin, plant_date, harvest_date, plant_price) VALUES ($1,$2,$3,$4,$5) RETURNING id;"

	err := db.DB.QueryRow(
		query,
		userId,
		coin,
		plantDate,
		harvestDate,
		plantPrice,
	).Scan(
		&plant.ID,
	)

	if err != nil {
		log.Fatalf("✡️  line 44 err %v", err)
	}

	return plant, nil
}

func GetPlant(userId int64, potId int) (t.Plant, error) {
	var plant t.Plant
	query := "SELECT * FROM farm.plant WHERE user_id=$1 AND id=$2;"

	err := db.DB.QueryRow(query, userId, potId).Scan(
		&plant.ID,
		&plant.UserID,
		&plant.Coin,
		&plant.PlantDate,
		&plant.HarvestDate,
		&plant.PlantPrice,
		&plant.HarvestPrice,
		&plant.Profit,
	)

	if err != nil {
		return plant, err
	}

	return plant, nil
}

func HarvestPlant(userId int64, plantId int, harvestPrice, profit float32) (t.Plant, error) {
	var plant t.Plant
	query := "UPDATE farm.plant SET harvest_price=$1, profit=$2 WHERE user_id=$3 AND id=$4 RETURNING *"

	err := db.DB.QueryRow(query, harvestPrice, profit, userId, plantId).Scan(
		&plant.ID,
		&plant.UserID,
		&plant.Coin,
		&plant.PlantDate,
		&plant.HarvestDate,
		&plant.PlantPrice,
		&plant.HarvestPrice,
		&plant.Profit,
	)

	if err != nil {
		return plant, err
	}

	return plant, nil
}

func GetPairPrice(t1, t2 t.Ticker) (float32, error) {
	url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s-%s/spot", t1, t2)
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var priceResp t.CoinbasePriceResponse

	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return 0, err
	}

	if price, err := strconv.ParseFloat(priceResp.Data.Amount, 32); err != nil {
		return 0, err
	} else {
		return float32(price), nil
	}
}
