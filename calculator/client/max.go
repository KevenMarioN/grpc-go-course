package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/KevenMarioN/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")
	stream, err := c.MAX(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MAXRequest{
		{Number: 20},
		{Number: 55},
		{Number: 5},
		{Number: 100},
	}

	waitc := make(chan struct{})
	go func() {
		for index := range reqs {
			log.Printf("Send number: %v\n", reqs[index])
			stream.Send(reqs[index])
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
