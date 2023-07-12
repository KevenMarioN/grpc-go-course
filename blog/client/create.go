package main

import (
	"context"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Keven",
		Title:    "Golang is special",
		Content:  "Learning gRPC",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Inexpected error: %v\n", err)
	}

	log.Printf("Blog hase been created: %v\n", res.Id)

	return res.Id
}
