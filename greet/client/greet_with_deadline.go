package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KevenMarioN/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Mário",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
