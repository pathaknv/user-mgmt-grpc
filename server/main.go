package main

import (
	pb "github.com/pathaknv/user-mgmt-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:5001"

type Server struct {
	pb.UserServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on port 5001: %v", err)
	}

	log.Printf("Listening on %s", addr)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
