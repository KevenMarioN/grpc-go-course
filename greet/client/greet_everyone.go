package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/KevenMarioN/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Mário"},
		{FirstName: "Pedro"},
		{FirstName: "João"},
	}

	waitc := make(chan struct{})
	go func() {
		for index := range reqs {
			log.Printf("Send request: %v\n", reqs[index])
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
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
