package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func doAVG(c pb.CalculatorServiceClient) {
	log.Println("doAVG was invoked")

	reqs := []*pb.AVGRequest{
		{Number: 5},
		{Number: 1},
		{Number: 25},
		{Number: 9},
	}

	stream, err := c.AVG(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAVG %v\n", err)
	}

	for index := range reqs {
		log.Printf("Sending req: %v\n", reqs[index])
		stream.Send(reqs[index])
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from doAVG: %v\n", err)
	}

	log.Printf("doAVG: %v\n", res.Result)
}
