package repository

//go:generate mockgen -source=repository.go -destination=./mock/mock_repository.go
type Repository interface {
	GetUrl(shortUrl uint64) (string, error) // TODO передавать по указателю или по значению?
	PostUrl(longUrl string) (uint64, error)
}
