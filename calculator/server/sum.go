package main

import (
	"context"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (res *pb.SumResponse, err error) {
	log.Printf("Sum function invoked with %v\n", req)
	sum := req.FirstNumber + req.SecondNumber
	return &pb.SumResponse{
		Result: sum,
	}, nil
}
