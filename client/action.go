package main

import (
	"context"
	pb "github.com/pathaknv/user-mgmt-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var grpcAddr string = "0.0.0.0:5001"

func createUser(params *pb.User) *pb.User {
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to connect gRPC server: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	user, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("Could not create user record: %v", err)
	}

	return user
}

func getUser(id *pb.Id) (*pb.User, error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to connect gRPC server: %v", err)
	}

	client := pb.NewUserServiceClient(conn)
	user, err := client.GetUser(context.Background(), id)

	return user, err
}
