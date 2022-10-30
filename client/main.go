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

	log.Print("Listening on ", addr)
	http.ListenAndServe(":5002", router)
}
