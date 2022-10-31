package main

import (
	"context"
	"github.com/pathaknv/user-mgmt-grpc/database"
	pb "github.com/pathaknv/user-mgmt-grpc/proto"
	"log"
)

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("CreateUser invoked on server side...")

	db := database.OpenDbConnection()
	defer db.Close()

	var id int32
	sqlQuery := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(sqlQuery, in.Name, in.Email).Scan(&id)

	return &pb.User{
		Id:    id,
		Name:  in.Name,
		Email: in.Email,
	}, err
}
