package tools

import (
	log "github.com/sirupsen/logrus"
)

// collections
type LoginDetails struct {
	AuthToken string
	Username string 
}

type CoinDetails struct {
	Coins int64
    Username string 
}

// create a db interface which will allow you to swap out dbs later
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// 1:05:36
