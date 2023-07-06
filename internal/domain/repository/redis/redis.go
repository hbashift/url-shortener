package redis

import (
	"context"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/errs"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

var notFoundError errs.NotFound
var connectionError errs.DatabaseConnectionError
var insertionError errs.InsertError
var alreadyExistsError errs.AlreadyExists

type redisDb struct {
	ctx      context.Context
	mainDB   *redis.Client
	uniqueDB *redis.Client
	id       uint64
}

type Config struct {
	Addr        string
	Pass        string
	DBNumMain   int
	DBNumUnique int
}

// TODO передавать по указателю или по значению?

func (r *redisDb) GetUrl(shortUrl uint64) (string, error) {
	val, err := r.mainDB.Get(r.ctx, strconv.FormatUint(shortUrl, 10)).Result()
	if err == redis.Nil {
		log.Printf("url not found: %v\n", err)
		notFoundError = fmt.Errorf("url not found: %w", err)

		return "", notFoundError
	} else if err != nil {
		log.Printf("could not connect to database: %v\n", err)
		connectionError = fmt.Errorf("could not connect to database: %v", err)

		return "", connectionError
	}

	return val, nil
}

// TODO передавать по указателю или по значению?

func (r *redisDb) PostUrl(longUrl string) (uint64, error) {
	set, err := r.uniqueDB.SetNX(r.ctx, longUrl, "", 0).Result()

	if set {
		r.id++
		err = r.mainDB.Set(r.ctx, strconv.FormatUint(r.id, 10), longUrl, 0).Err()
		if err != nil {
			log.Printf("could not insert into database%s: %v\n", r.uniqueDB, err)
			insertionError = fmt.Errorf("could not insert into database")

			return 0, insertionError
		}

		err = r.uniqueDB.Set(r.ctx, longUrl, "", 0).Err()
		if err != nil {
			log.Printf("could not insert into database%s: %v\n", r.uniqueDB, err)
			insertionError = fmt.Errorf("could not insert into database: %w", err)

			return 0, insertionError
		}

		err = r.mainDB.Set(r.ctx, "id", r.id, 0).Err()
		if err != nil {
			insertionError = fmt.Errorf("could not reset id value: %w", err)

			return 0, insertionError
		}

		return r.id, nil
	} else if !set {
		log.Printf("trying to insert already existing url: %v", longUrl)
		alreadyExistsError = fmt.Errorf("such url is already exists")

		return 0, alreadyExistsError
	} else {
		log.Printf("could not connect to database: %v\n", err)
		connectionError = fmt.Errorf("could not connect to database")

		return 0, connectionError
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

	idStr, err := mainDB.Get(ctx, "id").Result()
	if err == redis.Nil {
		idStr = "0"
		err = mainDB.Set(ctx, "id", idStr, 0).Err()

		if err != nil {
			log.Fatalf("could not set id key: %v\n", err)
		}
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Fatalf("could not parse id value: %v", err)
	}

	return &redisDb{
		ctx:      ctx,
		mainDB:   mainDB,
		uniqueDB: uniqueDB,
		id:       id,
	}
}
