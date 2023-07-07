package client

import (
	"context"
	"fmt"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
)

type shortenerClient struct {
	client pb.ShortenerClient
}

func (s *shortenerClient) PostUrl(ctx context.Context,
	longUrl *pb.LongUrl,
	opts ...grpc.CallOption) (*pb.ShortUrl, error) {

	resp, err := s.client.PostUrl(ctx, longUrl)
	if err != nil {
		return nil, fmt.Errorf("failure: %w", err)
	}

	return resp, nil
}

func (s *shortenerClient) GetUrl(ctx context.Context,
	shortUrl *pb.ShortUrl,
	opts ...grpc.CallOption) (*pb.LongUrl, error) {

	resp, err := s.client.GetUrl(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("failure: %w", err)
	}

	return resp, nil
}

func NewClient(conn *grpc.ClientConn) pb.ShortenerClient {
	return &shortenerClient{client: pb.NewShortenerClient(conn)}
}
