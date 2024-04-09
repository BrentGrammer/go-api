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

// func returns the db interface
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{} // note we set the mem addr to the variable (we're returning a pointer to the db interface)
    // the struct mockDB implements the db interface

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err // keep return sig in tact and return nil for db interface,
	}

	return &database, nil
}
