package redis

import (
	"context"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/redis/go-redis/v9"
)

type redisDb struct {
	ctx     context.Context
	mainDB  *redis.Client
	longDB  *redis.Client
	shortDB *redis.Client
}

type Config struct {
	Addr       string
	Pass       string
	DBNumMain  int
	DBNumLong  int
	DBNumShort int
}

func initRedis(cfg *Config) (*redis.Client, *redis.Client, context.Context) {
	return redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Pass,
			DB:       cfg.DBNumMain,
		}),
		redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Pass,
			DB:       cfg.DBNumLong,
		}),
		context.Background()
}

func NewRedis(cfg *Config) repository.Repository {
	mainDB, longDB, ctx := initRedis(cfg)

	return &redisDb{
		ctx:    ctx,
		mainDB: mainDB,
		longDB: longDB,
	}
}

// GetUrl - get method for url from redis database.
// If byLongUrl true - gets short url by long_url
// Else - gets by short_url
func (r *redisDb) GetUrl(url *model.Url, byLongUrl bool) (string, error) {
	var val string
	var err error

	if byLongUrl {
		val, err = r.longDB.Get(r.ctx, url.LongUrl).Result()
	} else {
		val, err = r.mainDB.Get(r.ctx, url.ShortUrl).Result()
	}

	if err == redis.Nil {

		return "", fmt.Errorf("cannot find url: %w", errs.ErrNotFound)
	} else if err != nil {

		return "", fmt.Errorf("could not connect to database: %v", errs.ErrDatabaseConnection)
	}

	return val, nil
}

// PostUrl -creates new record in redis database.
// If long_url exists - returns errs.ErrLongUrlExists.
// If short_url exists - return errs.ErrShortUrlExists.
// Else - returns errs.ErrDatabaseConnection
func (r *redisDb) PostUrl(url *model.Url) (string, error) {
	set, err := r.longDB.SetNX(r.ctx, url.LongUrl, url.ShortUrl, 0).Result()

	if err != nil {

		return "", fmt.Errorf("could not connect to database: %w", errs.ErrDatabaseConnection)

	} else {
		if set {
			err = r.longDB.Del(r.ctx, url.LongUrl).Err()
			if err != nil {
				return "", fmt.Errorf("could not delete: %w", err)
			}

			set, err = r.mainDB.SetNX(r.ctx, url.ShortUrl, url.LongUrl, 0).Result()
			if err != nil {

				return "", fmt.Errorf("could not insert into database: %w", errs.ErrInsertion)
			}

			if !set {
				return "", fmt.Errorf("such short url already exists: %w", errs.ErrShortUrlExists)
			}

			set, err = r.longDB.SetNX(r.ctx, url.LongUrl, url.ShortUrl, 0).Result()

			return url.ShortUrl, nil
		} else {

			return "", fmt.Errorf("such url is already exists: %w", errs.ErrLongUrlExists)
		}
	}
}
