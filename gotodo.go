package main

import (
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@(localhost:4211)/todolist?charset=utf8mb4&parseTime=True&loc=Local")

// TodoItem represents a todo item. Desc is the contents the user
// entered. Done shows whether the todo is completed or not.
type TodoItem struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Desc string `json:"desc"`
	Done bool   `json:"done"`
}

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
	defer db.Close()
	log.Info("GoTodo API Server")
	// Remove drop table if exists to persist database between runs
	db.Debug().DropTableIfExists(&TodoItem{})
	db.Debug().AutoMigrate(&TodoItem{})
	router := mux.NewRouter()
	router.HandleFunc("/isAlive", IsAlive).Methods("GET")
	http.ListenAndServe(":8000", router)
}
