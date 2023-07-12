package main

import (
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", req)

	k := int64(2)
	number := req.Number
	for number > 1 {
		if number%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			number /= k
		} else {
			k++
		}

	}
	return nil
}
