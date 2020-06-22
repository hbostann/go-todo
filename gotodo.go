package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

// ItemExists searches the DB for a ToDo item with a given id.
// Returns true if found.
func ItemExists(id int) bool {
	todo := &TodoItem{}
	result := db.First(todo, id)
	if result.Error != nil {
		log.Error("Todo item with given id not found in database", id)
		return false
	}
	return true
}

// CreateTodo is the API endpoint to add a new todo item.
func CreateTodo(response http.ResponseWriter, request *http.Request) {
	description := request.FormValue("description")
	log.WithField("description", description).Info("Adding todo item.")
	todo := &TodoItem{Desc: description, Done: false}
	db.Create(todo)
	result := db.Last(todo)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(result.Value)
}

// UpdateItem is the API endpoint to toggle the completion of a todo item
func UpdateItem(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("Given value is not a valid id", vars["id"])
		response.Header().Set("Content-Type", "application/json")
		io.WriteString(response, `{"updated":false, "error": "Invalid Request"}`)
		return
	}
	exists := ItemExists(id)
	if exists == false {
		response.Header().Set("Content-Type", "application/json")
		io.WriteString(response, `{"updated":false, "error": "Record Not Found}`)
		return
	}
	isDone, err := strconv.ParseBool(request.FormValue("done"))
	if err != nil {
		log.Error("Given todo id is not valid")
		response.Header().Set("Content-Type", "application/json")
		io.WriteString(response, `{"updated":false, "error": "Invalid Request"}`)
		return
	}
	todo := &TodoItem{}
	db.First(todo, id)
	todo.Done = isDone
	db.Save(todo)
	response.Header().Set("Content-Type", "application/json")
	io.WriteString(response, `{"updated": true}`)
}

// DeleteItem is the API endpoint for the deletion of a todo item.
func DeleteItem(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("Given value %v, is not a valid id", vars["id"])
		response.Header().Set("Content-Type", "application/json")
		io.WriteString(response, `{"deleted":false, "error": "Invalid Request"}`)
		return
	}
	exists := ItemExists(id)
	if exists == false {
		response.Header().Set("Content-Type", "application/json")
		io.WriteString(response, `{"deleted":false, "error": "Record Not Found}`)
		return
	}
	log.WithField("id", id).Info("Deleting todo")
	todo := &TodoItem{}
	db.First(todo, id)
	db.Delete(todo)
	response.Header().Set("Content-Type", "application/json")
	io.WriteString(response, `{"deleted": true}`)
}

// GetCompletedItems responds to the request with a json list of all the
// items which are marked as complete (Done == true).
func GetCompletedItems(response http.ResponseWriter, request *http.Request) {
	log.Info("Get Complete Items")
	var items []TodoItem
	db.Where("done = ?", true).Find(&items)
	log.Info("Complete Items: ", items)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(items)
}

// GetIncompleteItems responds to the request with a json list of all the
// items which are marked as incomplete (Done == false).
func GetIncompleteItems(response http.ResponseWriter, request *http.Request) {
	log.Info("Get Incomplete Items")
	var items []TodoItem
	db.Where("done = ?", false).Find(&items)
	log.Info("Incomplete Items: ", items)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(items)
}

// GetItems responds to the request with a json list of all existings todos
// in the database.
func GetItems(response http.ResponseWriter, request *http.Request) {
	log.Info("Get All Items")
	var items []TodoItem
	db.Find(&items)
	log.Info("All Items: ", items)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(items)
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
