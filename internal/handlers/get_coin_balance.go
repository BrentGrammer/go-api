package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BrentGrammer/goapi/api" // bring in entities types
	"github.com/BrentGrammer/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema" // gorilla used to decode params
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// this happens after auth middleware passed and we need to decode user params
	var params = api.CoinBalanceParams{} // struct is defined in api/api.go and represents the url parameters for this request
    
	// use gorilla package to decode
	var decoder *schema.Decoder = schema.NewDecoder() 
	var err error
    // this grabs and decodes the parameters(username) from the url and set it to the values in the params struct
	err = decoder.Decode(&params,  r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil { // will be nil if not found in map
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins, // details from getusercoins is a pointer so we dereference to get the value
		Code: http.StatusOK,
	}

	// write response to response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}