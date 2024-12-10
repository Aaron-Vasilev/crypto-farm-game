package consts

import t "crypto-farm/src/types"

var (
	UserID = "UserID"

	Plants = map[t.Ticker]t.Plant{
		t.DOGE: {
			Coin:    t.DOGE,
			Cost:    5,
			Exp:     10,
			Minutes: 1,
		},
	}
)
