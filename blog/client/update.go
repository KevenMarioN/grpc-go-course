package main

import (
	"context"
	"log"

	pb "github.com/KevenMarioN/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Clement",
		Title:    "A new title",
		Content:  "Content Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}
	log.Println("Blog was updated!")
}
