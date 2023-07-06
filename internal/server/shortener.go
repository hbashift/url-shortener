package server

import (
	"context"
	"errors"
	"github.com/hbashift/url-shortener/internal/domain/errs"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ShortenerServer struct {
	shortener *service.ShortenerService
	pb.UnimplementedShortenerServer
}

var notFoundError errs.NotFound

var insertionError errs.InsertError

var conncectionError errs.DatabaseConnectionError

var migrationError errs.DatabaseMigrationError

var alreadyExistsError errs.AlreadyExists

func (s *ShortenerServer) PostUrl(ctx context.Context, url *pb.LongUrl) (*pb.ShortUrl, error) {
	shortUrl, err := s.shortener.PostUrl(url.GetLongUrl())
	/*var notFoundError errs.NotFound
	var alreadyExistsError errs.AlreadyExists
	var connectionError errs.DatabaseConnectionError
	var migrationError errs.DatabaseMigrationError*/
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	if err != nil {
		if errors.As(err, &alreadyExistsError) {
			return nil, status.Errorf(codes.AlreadyExists, alreadyExistsError.Error())
		} else if errors.As(err, &conncectionError) {
			return nil, status.Errorf(codes.Internal, conncectionError.Error())
		} else if errors.As(err, &migrationError) {
			return nil, status.Errorf(codes.Internal, migrationError.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return shortUrl, nil
}

func (s *ShortenerServer) GetUrl(ctx context.Context, url *pb.ShortUrl) (*pb.LongUrl, error) {
	if len([]rune(url.GetShortUrl())) > 10 {
		return nil, status.Errorf(codes.InvalidArgument, "short url length must be 10")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	longUrl, err := s.shortener.GetUrl(url.GetShortUrl())

	if err != nil {
		if errors.As(err, &notFoundError) {
			return nil, status.Errorf(codes.InvalidArgument, notFoundError.Error())
		} else if errors.As(err, &conncectionError) {
			return nil, status.Errorf(codes.Internal, conncectionError.Error())
		} else if errors.As(err, &migrationError) {
			return nil, status.Errorf(codes.Internal, migrationError.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return longUrl, nil
}

func NewShortenerServer(s *service.ShortenerService) *ShortenerServer {
	return &ShortenerServer{shortener: s}
}
