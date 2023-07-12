package main

import (
	"io"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func (s *Server) MAX(stream pb.CalculatorService_MAXServer) error {
	log.Println("MAX function invoked")

	max := int64(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number
			log.Println(req.Number)
			err = stream.Send(&pb.MAXResponse{
				Result: req.Number,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
