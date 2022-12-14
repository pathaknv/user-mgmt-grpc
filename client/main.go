package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var addr string = "0.0.0.0:5002"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")

	log.Print("Listening on ", addr)
	http.ListenAndServe(":5002", router)
}
