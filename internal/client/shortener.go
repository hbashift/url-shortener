package client

import (
	"context"
	"fmt"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Shortener struct {
	client pb.ShortenerClient
}

func (s *Shortener) PostUrl(ctx context.Context, longUrl *pb.LongUrl, opts ...grpc.CallOption) (*pb.ShortUrl, error) {
	resp, err := s.client.PostUrl(ctx, longUrl)
	if err != nil {
		return nil, fmt.Errorf("failure: %w", err)
	}

	return resp, nil
}

func (s *Shortener) GetUrl(ctx context.Context, shortUrl *pb.ShortUrl, opts ...grpc.CallOption) (*pb.LongUrl, error) {
	resp, err := s.client.GetUrl(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("failure: %w", err)
	}

	return resp, nil
}

func RunClient(URL string) pb.ShortenerClient {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewShortenerClient(conn)
	return client
}
