package model

// Url - struct for PostgreSQL database
type Url struct {
	ShortUrl string `gorm:"primarykey"`
	LongUrl  string `gorm:"unique; not null"`
}

// ShortUrl - DTO for gRPC gateway HTTP server
type ShortUrl struct {
	ShortUrl string `json:"shortUrl"`
}

// LongUrl - DTO for gRPC gateway HTTP server
type LongUrl struct {
	LongUrl string `json:"longUrl"`
}
