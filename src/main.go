package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.con/nicolas/golang_api/src/api"
)

func main() {
	var port string = "8000"
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	fmt.Printf("Server running ar port %s", port)
	http.ListenAndServe(":"+port, router)
}
