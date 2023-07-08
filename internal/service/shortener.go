package service

import (
	"errors"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/hbashift/url-shortener/internal/util/encoder"
	"math/rand"
)

/*
ShortenerService - struct for internal logic of the whole service.
db repository.Repository - interface for db layer of the service
*/
type ShortenerService struct {
	db repository.Repository
}

/*
NewShortenerService - ShortenerService initialization
*/
func NewShortenerService(db repository.Repository) *ShortenerService {
	return &ShortenerService{db: db}
}

// PostUrl - encodes input URL with random generated id from 10based into 63based
func (s *ShortenerService) PostUrl(longUrl string) (string, error) {
	id := rand.Uint64()
	shortUrl := encoder.EncodeUrl(id, false)
	url := model.Url{
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	}

	shortUrl, err := s.db.PostUrl(&url)
	if err != nil {
		for errors.Is(err, errs.ErrShortUrlExists) {
			id = rand.Uint64()
			shortUrl = encoder.EncodeUrl(id, true)

			url.ShortUrl = shortUrl
			shortUrl, err = s.db.PostUrl(&url)

			if err == nil {
				break
			} else if !errors.Is(err, errs.ErrShortUrlExists) {
				return "", fmt.Errorf("db error: %w", err)
			}
		}

		if errors.Is(err, errs.ErrLongUrlExists) {
			shortUrl, err = s.db.GetUrl(&url, true)

			if err != nil {
				return "", fmt.Errorf("such long url not found: %w", errs.ErrNotFound)
			}
		}
	}

	return shortUrl, nil
}

// GetUrl - get method for original URL. shortUrl is an input url
func (s *ShortenerService) GetUrl(shortUrl string) (string, error) {
	url := model.Url{ShortUrl: shortUrl}

	longUrl, err := s.db.GetUrl(&url, false)
	if err != nil {

		return "", fmt.Errorf("could not get original longUrl: %w", err)
	}

	return longUrl, nil
}
