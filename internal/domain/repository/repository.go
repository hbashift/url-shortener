package repository

import "github.com/hbashift/url-shortener/internal/domain/repository/model"

//go:generate mockgen -source=repository.go -destination=./mock/mock_repository.go

// Repository - interface for different database realizations
type Repository interface {
	GetUrl(url *model.Url, byLongUrl bool) (string, error)
	PostUrl(url *model.Url) (string, error)
}
