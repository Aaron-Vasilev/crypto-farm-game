package controller

import (
	"crypto-farm/src/consts"
	"crypto-farm/src/db"
	t "crypto-farm/src/types"
	"database/sql"
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

func PlantCoin(
	userId int64,
	potId int,
	coin t.Ticker,
	plantTime time.Time,
	harvestTime time.Time,
	plantPrice float32,
) (t.Pot, error) {
	pot := t.Pot{
		UserID:      userId,
		Coin:        coin,
		PlantTime:   plantTime,
		HarvestTime: harvestTime,
		PlantPrice:  plantPrice,
	}
	query := "UPDATE farm.pot SET coin=$1, plant_time=$2, harvest_time=$3, plant_price=$4 WHERE user_id=$5 AND id=$6 RETURNING id;"

	err := db.DB.QueryRow(
		query,
		coin,
		plantTime,
		harvestTime,
		plantPrice,
		userId,
		potId,
	).Scan(
		&pot.ID,
	)

	if err != nil {
		log.Fatalf("✡️  line 44 err %v", err)
	}

	return pot, nil
}

// func GetPlant(userId int64, potId int) (t.Plant, error) {
// 	var plant t.Plant
// 	query := "SELECT * FROM farm.plant WHERE user_id=$1 AND id=$2;"

// 	err := db.DB.QueryRow(query, userId, potId).Scan(
// 		&plant.ID,
// 		&plant.UserID,
// 		&plant.Coin,
// 		&plant.PlantDate,
// 		&plant.HarvestDate,
// 		&plant.PlantPrice,
// 		&plant.HarvestPrice,
// 		&plant.Profit,
// 	)

// 	if err != nil {
// 		return plant, err
// 	}

// 	return plant, nil
// }

// func HarvestPlant(userId int64, plantId int, harvestPrice, profit float32) (t.Plant, error) {
// 	var plant t.Plant
// 	query := "UPDATE farm.plant SET harvest_price=$1, profit=$2 WHERE user_id=$3 AND id=$4 RETURNING *"

// 	err := db.DB.QueryRow(query, harvestPrice, profit, userId, plantId).Scan(
// 		&plant.ID,
// 		&plant.UserID,
// 		&plant.Coin,
// 		&plant.PlantDate,
// 		&plant.HarvestDate,
// 		&plant.PlantPrice,
// 		&plant.HarvestPrice,
// 		&plant.Profit,
// 	)

// 	if err != nil {
// 		return plant, err
// 	}

// 	return plant, nil
// }

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

func GetPotsWithPlants(userId int64) ([]t.PotWithPlant, error) {
	var pp []t.PotWithPlant
	pots, err := GetPots(userId)

	if err != nil {
		return pp, err
	}

	for _, pot := range pots {
		pp = append(pp, CombinePlantAndCoin(pot))
	}

	return pp, err
}

func CombinePlantAndCoin(pot t.Pot) t.PotWithPlant {
	return t.PotWithPlant{
		Pot:   pot,
		Plant: consts.Plants[pot.Coin],
	}
}

func GetPots(userId int64) ([]t.Pot, error) {
	var pots []t.Pot
	query := "SELECT * FROM farm.pot WHERE user_id=$1;"

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return pots, err
	}

	for rows.Next() {
		var p t.Pot
		var coin sql.NullString
		var plantDate sql.NullTime
		var harvestDate sql.NullTime
		var plantPrice sql.NullFloat64
		var harvestPrice sql.NullFloat64

		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&coin,
			&plantDate,
			&harvestDate,
			&plantPrice,
			&harvestPrice,
		)

		if err != nil {
			return pots, err
		}

		if coin.Valid {
			p.Coin = t.Ticker(coin.String)
			p.PlantTime = plantDate.Time
			p.HarvestTime = harvestDate.Time
			p.PlantPrice = float32(plantPrice.Float64)

			if harvestPrice.Valid {
				p.HarvestPrice = float32(harvestPrice.Float64)
			}
		}

		pots = append(pots, p)
	}
	defer rows.Close()

	return pots, nil
}

func GetPotById(userId int64, potId int) (t.Pot, error) {
	var pot t.Pot
	var coin sql.NullString
	var plantDate sql.NullTime
	var harvestDate sql.NullTime
	var plantPrice sql.NullFloat64
	var harvestPrice sql.NullFloat64
	query := "SELECT * FROM farm.pot WHERE user_id=$1 AND id=$2;"

	row := db.DB.QueryRow(query, userId, potId)

	if row.Err() != nil {
		return pot, row.Err()
	}

	row.Scan(
		&pot.ID,
		&pot.UserID,
		&coin,
		&plantDate,
		&harvestDate,
		&plantPrice,
		&harvestPrice,
	)

	if coin.Valid {
		pot.Coin = t.Ticker(coin.String)
		pot.PlantTime = plantDate.Time
		pot.HarvestTime = harvestDate.Time
		pot.PlantPrice = float32(plantPrice.Float64)

		if harvestPrice.Valid {
			pot.HarvestPrice = float32(harvestPrice.Float64)
		}
	}

	return pot, nil
}

func UpdateHarvestPrice(userId int64, potId int, price float32) (t.Pot, error) {
	var pot t.Pot
	query := "UPDATE farm.pot SET harvest_price=$1 WHERE user_id=$2 AND id=$3 RETURNING *;"

	row := db.DB.QueryRow(query, userId, potId)

	if row.Err() != nil {
		return pot, row.Err()
	}

	row.Scan(
		&pot.ID,
		&pot.UserID,
		&pot.Coin,
		&pot.PlantTime,
		&pot.HarvestTime,
		&pot.PlantPrice,
		&pot.HarvestPrice,
	)

	return pot, nil
}
