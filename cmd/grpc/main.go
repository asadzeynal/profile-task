package main

import (
	"log"
	"net"

	pb "github.com/asadzeynal/profile-task/gen/profile/v1"
	"github.com/asadzeynal/profile-task/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := profile.NewClient()
	server := profile.NewServer(client)

	// TODO: Implement graceful shutdown
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
