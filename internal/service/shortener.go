package service

import (
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/util/encoder"
	shortener "github.com/hbashift/url-shortener/proto"
	"log"
)

type ShortenerService struct {
	db repository.Repository
}

func NewShortenerService(db repository.Repository) *ShortenerService {
	return &ShortenerService{db: db}
}

func (s *ShortenerService) PostUrl(longUrl string) (*shortener.ShortUrl, error) {
	id, err := s.db.PostUrl(longUrl)
	if err != nil {
		log.Printf("could not post url: %v", err)
		return nil, err
	}

	res, err := encoder.EncodeUrl(id)
	if err != nil {
		log.Printf("could not encode url: %v", err)
		return nil, err
	}

	return &shortener.ShortUrl{ShortUrl: res}, nil
}

func (s *ShortenerService) GetUrl(shortUrl string) (*shortener.LongUrl, error) {
	id, err := encoder.DecryptUrl(shortUrl)
	if err != nil {
		log.Printf("could not decrypt url: %v", err)
		return nil, err
	}

	longUrl, err := s.db.GetUrl(id)
	if err != nil {
		log.Printf("could not get original url: %v", err)
		return nil, err
	}

	return &shortener.LongUrl{LongUrl: longUrl}, nil
}
