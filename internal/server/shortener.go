package server

import (
	"context"
	"errors"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type shortenerServer struct {
	shortener service.ShortenerService
	pb.UnimplementedShortenerServer
}

func (s *shortenerServer) PostUrl(ctx context.Context, url *pb.LongUrl) (*pb.ShortUrl, error) {
	if len([]rune(url.GetLongUrl())) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "url length must > 0")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	shortUrl, err := s.shortener.PostUrl(url.GetLongUrl())

	if err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := pb.ShortUrl{ShortUrl: shortUrl}

	return &res, nil
}

func (s *shortenerServer) GetUrl(ctx context.Context, url *pb.ShortUrl) (*pb.LongUrl, error) {
	if len([]rune(url.GetShortUrl())) != 10 {
		return nil, status.Errorf(codes.InvalidArgument, "short url length must be 10")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	longUrl, err := s.shortener.GetUrl(url.GetShortUrl())

	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := pb.LongUrl{LongUrl: longUrl}

	return &res, nil
}

func NewShortenerServer(s service.ShortenerService) pb.ShortenerServer {
	return &shortenerServer{shortener: s}
}
