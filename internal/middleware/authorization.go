package middleware

import (
	"errors"
	"net/http"

	"github.com/BrentGrammer/goapi/api"
	"github.com/BrentGrammer/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

// Middleware must follow a signature - must take in and return a http.Handler
func Authorization(next http.Handler) http.Handler {
	// create a handler using HandlerFunc which takes a function that has a response writer and pointer to the request:
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError) // sends error to client using our error handler
			return // exit the function after logging and sending error
		}

		// instantiate database to proceed if no error
		var database *tools.DatabaseInterface // use a pointer
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w) // return 500 using our handler
			return // exit
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}		

		next.ServeHTTP(w, r) // call next with writer and response passed in at end of middleware 
	})
}
