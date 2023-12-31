package main

import (
	"context"
	"fmt"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

// main - runs simple gRPC client.
// Run gRPC server first
func main() {
	log.Print("starting")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewShortenerClient(conn)
	request := &pb.LongUrl{
		LongUrl: "https://localhost:6969/test_url/",
	}
	response, err := client.PostUrl(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.GetShortUrl())
}
