package main

import (
	"context"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	log.Printf("Greet function was invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
