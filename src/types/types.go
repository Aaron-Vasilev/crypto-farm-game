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
	ID      int
	UserID  int64
	PlantID *int
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

// NOTE not sure that I need it
type Plant struct {
	ID           int       `json:"id,omitempty"`
	UserID       int64     `json:"userId,omitempty"`
	Coin         Ticker    `json:"coin"`
	PlantDate    time.Time `json:"plantDate,omitempty"`
	HarvestDate  time.Time `json:"harvestDate,omitempty"`
	PlantPrice   float32   `json:"plantPrice,omitempty"`
	HarvestPrice *float32  `json:"harvestPrice,omitempty"`
	Profit       *float32  `json:"profit,omitempty"`
}

type CoinbasePriceResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}
