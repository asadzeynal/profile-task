package main

import (
	"log"
	"net"

	"github.com/asadzeynal/profile-task/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := NewClient()
	server := NewServer(client)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterProfileServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting grpc server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
