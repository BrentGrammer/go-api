package api

import (
	"encoding/json"
	"net/http"
)

//  struct to define parameters in the URL for the endpoint
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

// make a reusable helper to write an error response to the http writer
// this is a general error writer and we can pass in specific errors we want to send back
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// we want some different types of errors so we wrap the above func for use (note the capital letters for exportable)
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unepxected Error Occured.", http.StatusInternalServerError)
	} 
)