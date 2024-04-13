package main

// the chi package is a flexible package to use for web development
// you can install these standard packages using `go mod tidy`
// import our own module from the internal handlers folder
import (
  "fmt"
  "net/http"

  "github.com/go-chi/chi"
  "github.com/BrentGrammer/goapi/internal/handlers" // used as handlers.{method}
  log "github.com/sirupsen/logrus" // alias to log
)

func main() {
  // set up logger so you get the file and line number
  log.SetReportCaller(true)

  // create a Mux router pointer
  var r *chi.Mux = chi.NewRouter()
  // pass the router into the handler function
  handlers.Handler(r)

  const PORT = 8000

  fmt.Printf("Starting GO API service to listen on port %d...", PORT)
  
  // NOTE: you need to use 0.0.0.0 to specify accepting incoming requests inside the docker container and not "localhost"
  // You could also just use ":8000" without specifying localhost or 0.0.0.0
  var serverUrl = fmt.Sprintf("0.0.0.0:%d", PORT) // use Sprintf to return a string and not print to stdout

  // start the server - pass in the base location and router
  err := http.ListenAndServe(serverUrl, r) // note this is a blocking call
  if err != nil {
    log.Error(err)
  }
}