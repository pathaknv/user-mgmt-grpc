package main

import (
	"context"
	"database/sql"
	"errors"
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

func (s *Server) GetUser(ctx context.Context, in *pb.Id) (*pb.User, error) {
	log.Printf("GetUser invoked on server side...")

	db := database.OpenDbConnection()
	defer db.Close()

	var id int32
	var name string
	var email string

	sqlQuery := "SELECT * FROM users WHERE id = $1"
	err := db.QueryRow(sqlQuery, in.Id).Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.User{}, errors.New("No record found!")
		} else {
			log.Fatalf("Query failed: %v", err)
		}
	}

	user := &pb.User{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return user, err
}

func (s *Server) UpdateUser(ctx context.Context, userParams *pb.User) (*pb.User, error) {
	log.Printf("UpdateUser invoked on server side...")

	db := database.OpenDbConnection()
	defer db.Close()

	sqlQuery := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	result, err := db.Exec(sqlQuery, userParams.Name, userParams.Email, userParams.Id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	} else if rowsAffected == 0 {
		return &pb.User{}, errors.New("Failed to update record!")
	}

	return &pb.User{}, nil
}
