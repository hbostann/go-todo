package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// IsAlive is the api function to check whether the backend is
// running as expected.
func IsAlive(response http.ResponseWriter, request *http.Request) {
	log.Info("API is alive")
	response.Header().Set("Content-Type", "application/json")
	io.WriteString(response, `{"isAlive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetReportCaller(false)
}

func main() {
	log.Info("GoTodo API Server")
	router := mux.NewRouter()
	router.HandleFunc("/isAlive", IsAlive).Methods("GET")
	http.ListenAndServe(":8000", router)
}
