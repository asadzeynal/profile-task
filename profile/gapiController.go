package profile

import (
	"context"
	"fmt"
	"log"

	pb "github.com/asadzeynal/profile-task/gen/profile/v1"
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
