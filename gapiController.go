package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asadzeynal/profile-task/pb"
)

type Server struct {
	pb.UnimplementedProfileServiceServer
	client *Client
}

func NewServer(c *Client) *Server {
	return &Server{client: c}
}

func (s *Server) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	companyData, err := s.client.FetchDataByInn(req.GetInn())
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	return &pb.GetProfileResponse{Inn: companyData.Inn,
		Kpp:           companyData.Kpp,
		OwnerFullName: companyData.OwnerFullName,
		CompanyName:   companyData.CompanyName,
	}, nil
}
