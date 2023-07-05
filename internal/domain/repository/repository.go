package repository

type Repository interface {
	GetUrl(shortUrl uint64) (string, error) // TODO передавать по указателю или по значению?
	PostUrl(longUrl string) (uint64, error)
}
