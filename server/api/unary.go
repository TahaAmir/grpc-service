package api

import (
	"context"

	pb "github.com/TahaAmir/grpc-service/proto"
)

type HelloServer struct {
	pb.GreetServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "Hello"}, nil
}
