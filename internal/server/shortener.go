package server

import (
	"context"
	"github.com/hbashift/url-shortener/internal/service"
	shortener "github.com/hbashift/url-shortener/proto"
)

type ShortenerServer struct {
	shortener *service.ShortenerService
	shortener.UnimplementedShortenerServer
}

func (s *ShortenerServer) PostUrl(ctx context.Context, url *shortener.LongUrl) (*shortener.ShortUrl, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShortenerServer) GetUrl(ctx context.Context, url *shortener.ShortUrl) (*shortener.LongUrl, error) {
	//TODO implement me
	panic("implement me")
}
