package redis

import (
	"context"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/redis/go-redis/v9"
	"log"
)

type redisDb struct {
	ctx      context.Context
	mainDB   *redis.Client
	uniqueDB *redis.Client
}

type Config struct {
	Addr        string
	Pass        string
	DBNumMain   int
	DBNumUnique int
}

// TODO Внутри докера можно обращаться как http://redis:PORT

// TODO передавать по указателю или по значению?

func (r *redisDb) GetUrl(url *model.Url) (string, error) {
	val, err := r.mainDB.Get(r.ctx, url.ShortUrl).Result()
	if err == redis.Nil {
		log.Printf("url not found: %v\n", err)

		return "", fmt.Errorf("cannot find url: %w", errs.ErrNotFound)
	} else if err != nil {
		log.Printf("could not connect to database: %v\n", err)

		return "", fmt.Errorf("could not connect to database: %v", errs.ErrDatabaseConnection)
	}

	return val, nil
}

// TODO передавать по указателю или по значению?

func (r *redisDb) PostUrl(url *model.Url) (string, error) {
	set, err := r.uniqueDB.SetNX(r.ctx, url.LongUrl, "", 0).Result()

	if set {
		set, err = r.mainDB.SetNX(r.ctx, url.ShortUrl, url.LongUrl, 0).Result()
		if err != nil {
			log.Printf("could not insert into database%s: %v\n", "redis_0", errs.ErrInsertion)

			return "", fmt.Errorf("could not insert into database: %w", errs.ErrInsertion)
		}

		if !set {
			/*err = r.uniqueDB.Set(r.ctx, url.LongUrl, "", 0).Err()
			if err != nil {
				log.Printf("could not insert into database%s: %v\n", "redis_1", errs.ErrInsertion)

				return "", fmt.Errorf("could not insert into database: %w", errs.ErrInsertion)
			}*/

			return "", fmt.Errorf("such short url already exists: %w", errs.ErrShortUrlExists)

		}

		return url.ShortUrl, nil
	} else if !set {
		log.Printf("trying to insert already existing url %v: %v", url, errs.ErrAlreadyExists)

		return "", fmt.Errorf("such url is already exists: %w", errs.ErrLongUrlExists)
	} else {
		log.Printf("could not connect to database: %v\n", errs.ErrDatabaseConnection)

		return "", fmt.Errorf("could not connect to database: %w", errs.ErrDatabaseConnection)
	}
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
			DB:       cfg.DBNumUnique,
		}),
		context.Background()
}

func NewRedis(cfg *Config) repository.Repository {
	mainDB, uniqueDB, ctx := initRedis(cfg)

	return &redisDb{
		ctx:      ctx,
		mainDB:   mainDB,
		uniqueDB: uniqueDB,
	}
}
