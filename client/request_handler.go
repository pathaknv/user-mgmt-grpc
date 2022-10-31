package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	pb "github.com/pathaknv/user-mgmt-grpc/proto"
	"log"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request: POST /users")

	var user pb.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser(&user)

	fmt.Fprintf(w, "New user created!")
	log.Printf("Completed request POST /users")
	log.Printf("----------------------------------------------")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request: GET /users/:id")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse ID")
	} else {
		idStruct := pb.Id{Id: int32(id)}
		user, err := getUser(&idStruct)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprint(err))
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}

	log.Printf("Completed request GET /users/:id")
	log.Printf("----------------------------------------------")
}
