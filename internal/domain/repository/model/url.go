package model

type Url struct {
	ShortUrl string `gorm:"primarykey"`
	LongUrl  string `gorm:"unique; not null"`
}

type ShortUrl struct {
	ShortUrl string `json:"shortUrl"`
}

type LongUrl struct {
	LongUrl string `json:"longUrl"`
}
