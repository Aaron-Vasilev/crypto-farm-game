package types

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  *string
	Username  *string
	Balance   float32
}

type Pot struct {
	ID           int
	UserID       int64
	Coin         Ticker
	PlantTime    time.Time
	HarvestTime  time.Time
	PlantPrice   float32
	HarvestPrice float32
}

type Ticker string

const (
	BTC  Ticker = "BTC"
	TON  Ticker = "TON"
	ETH  Ticker = "ETH"
	DOGE Ticker = "DOGE"
	SOL  Ticker = "SOL"
	NEAR Ticker = "NEAR"
	USD  Ticker = "USD"
)

type Plant struct {
	Coin    Ticker
	Cost    int
	Exp     int
	Minutes int
}

type CoinbasePriceResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

type PotWithPlant struct {
	Plant
	Pot
}
