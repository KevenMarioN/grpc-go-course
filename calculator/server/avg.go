package main

import (
	"io"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func (s *Server) AVG(stream pb.CalculatorService_AVGServer) error {
	log.Println("AVG function was invoked")

	var (
		res     int64
		counter int64
	)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AVGResponse{
				Result: float64(res) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += req.Number
		counter++
	}
}
