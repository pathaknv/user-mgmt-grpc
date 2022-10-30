package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/pathaknv/user-mgmt-grpc/proto"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request: POST /users")

	var user pb.UserParams
	json.NewDecoder(r.Body).Decode(&user)

	createUser(&user)

	fmt.Fprintf(w, "New user created!")
	log.Printf("Completed request POST /users")
	log.Printf("----------------------------------------------")
}
