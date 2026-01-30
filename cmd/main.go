package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salsapunk/API-REST/db"
	"github.com/salsapunk/API-REST/handlers"
	"github.com/salsapunk/API-REST/repository"
	"github.com/salsapunk/API-REST/response"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Response{
		Message: "pong",
		Status:  200,
	})
}

func main() {
	mux := mux.NewRouter()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	NewTaskRepository := repository.NewTaskRepository(dbConnection)
	NewTaskHandler := handlers.NewTaskHandler(&NewTaskRepository)

	mux.HandleFunc("/ping", healthHandler)
	mux.HandleFunc("/GET/tasks", NewTaskHandler.ListAll)
	mux.HandleFunc("/POST/task", NewTaskHandler.Create)
	// mux.HandleFunc("/PUT/task/:id", NewTaskHandler.Edit)

	http.Handle("/", mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
