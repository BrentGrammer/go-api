package tools

import (
	"time"
)

// initialize mock struct to impl db interface
// below we attach methods with receivers for the impl
type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex"
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason"
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie"
	}
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex"
	},
	"jason": {
		Coins: 200,
		Username: "jason"
	},
	"marie": {
		Coins: 300,
		Username: "marie"
	}
}

// implement interface and attach methods to mockDB with receivers
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//simulate call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//simulate call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData // why do we return the mem addr/pointer here?
}

// setup does nothing
func (d *mockDB) SetupDatabase() error {
	return nil
}