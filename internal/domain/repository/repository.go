package repository

import "github.com/hbashift/url-shortener/internal/domain/repository/model"

//go:generate mockgen -source=repository.go -destination=./mock/mock_repository.go
type Repository interface {
	GetUrl(url *model.Url) (string, error) // TODO передавать по указателю или по значению?
	PostUrl(url *model.Url) (string, error)
}
