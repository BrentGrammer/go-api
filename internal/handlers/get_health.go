package handlers

import (
    "fmt"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    fmt.Fprintf(w, `{"message": "Successfully hit root route"}`)
}