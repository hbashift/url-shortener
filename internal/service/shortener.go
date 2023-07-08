package service

import (
	"errors"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/hbashift/url-shortener/internal/util/encoder"
	"log"
	"math/rand"
)

type ShortenerService struct {
	db repository.Repository
}

func NewShortenerService(db repository.Repository) *ShortenerService {
	return &ShortenerService{db: db}
}

func (s *ShortenerService) PostUrl(longUrl string) (string, error) {
	id := rand.Uint64()
	shortUrl := encoder.EncodeUrl(id, false)
	url := model.Url{
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	}

	shortUrl, err := s.db.PostUrl(&url)
	if err != nil {
		log.Printf("could not post url: %v", err)

		if errors.Is(err, errs.ErrShortUrlExists) {
			for err != nil {
				id = rand.Uint64()
				shortUrl = encoder.EncodeUrl(id, true)

				url.ShortUrl = shortUrl
				shortUrl, err = s.db.PostUrl(&url)
			}

		} else if errors.Is(err, errs.ErrLongUrlExists) {
			return "", fmt.Errorf("such long url already exists: %w", errs.ErrAlreadyExists)
		}
	}

	return shortUrl, nil
}

func (s *ShortenerService) GetUrl(shortUrl string) (string, error) {
	url := model.Url{ShortUrl: shortUrl}

	longUrl, err := s.db.GetUrl(&url)
	if err != nil {

		return "", fmt.Errorf("could not get original longUrl: %w", err)
	}

	return longUrl, nil
}
