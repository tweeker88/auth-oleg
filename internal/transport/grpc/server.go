package grpc

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/tweeker88/auth-oleg/pkg/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	user_v1.UnimplementedUserV1Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(_ context.Context, _ *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	return &user_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *Server) Get(_ context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	now := timestamppb.New(time.Now())

	return &user_v1.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      user_v1.Role_ADMIN,
		UpdatedAt: now,
		CreatedAt: now,
	}, nil
}

func (s *Server) Update(_ context.Context, _ *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(_ context.Context, _ *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
