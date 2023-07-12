package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KevenMarioN/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("DoLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Keven"},
		{FirstName: "Mário"},
		{FirstName: "Thales"},
		{FirstName: "Vitória"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreetL %v\n", err)
	}

	for index := range reqs {
		log.Printf("Sending req: %v\n", reqs[index])
		stream.Send(reqs[index])
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
