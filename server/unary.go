package main

import (
	"context"

	pb "github.com/TahaAmir/grpc-service/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "Hello from server"}, nil
}
