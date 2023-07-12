package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/KevenMarioN/grpc-go-course/blog/proto"
)

const addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "ANOTHINGEXIST")
	updateBlog(c, id)
	listBlog(c)

	id = createBlog(c)
	deleteBlog(c, id)
}
