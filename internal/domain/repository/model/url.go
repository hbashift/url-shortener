package model

type Url struct {
	ID  uint64 `gorm:"type:bigserial; primarykey"`
	Url string `gorm:"unique; not null"`
}

type ShortUrl struct {
	ShortUrl string `json:"shortUrl"`
}

type LongUrl struct {
	LongUrl string `json:"longUrl"`
}
