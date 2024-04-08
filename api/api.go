package api

import (
	"encoding/json"
	"net/http"
)

//  struct to define parameters endpoint will take
type CoinBalanceParams struct {
	Username string
}

// response struct
type CoinBalanceResponse struct {
	Code int // response code
	Balance int64
}

// Error struct
type Error struct {
	Code int
	Message string
}