package redis

import (
	"context"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

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

// TODO Внутри докера можно обращаться как http://redis:PORT

// TODO передавать по указателю или по значению?

func (r *redisDb) GetUrl(shortUrl uint64) (string, error) {
	val, err := r.mainDB.Get(r.ctx, strconv.FormatUint(shortUrl, 10)).Result()
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

func (r *redisDb) PostUrl(longUrl string) (uint64, error) {
	set, err := r.uniqueDB.SetNX(r.ctx, longUrl, "", 0).Result()

	if set {
		r.id++
		err = r.mainDB.Set(r.ctx, strconv.FormatUint(r.id, 10), longUrl, 0).Err()
		if err != nil {
			log.Printf("could not insert into database%s: %v\n", "redis_0", errs.ErrInsertion)

			return 0, fmt.Errorf("could not insert into database: %w", errs.ErrInsertion)
		}

		err = r.uniqueDB.Set(r.ctx, longUrl, "", 0).Err()
		if err != nil {
			log.Printf("could not insert into database%s: %v\n", "redis_1", errs.ErrInsertion)

			return 0, fmt.Errorf("could not insert into database: %w", errs.ErrInsertion)
		}

		err = r.mainDB.Set(r.ctx, "id", r.id, 0).Err()
		if err != nil {

			return 0, fmt.Errorf("could not reset id value: %w", errs.ErrInsertion)
		}

		return r.id, nil
	} else if !set {
		log.Printf("trying to insert already existing url %v: %v", longUrl, errs.ErrAlreadyExists)

		return 0, fmt.Errorf("such url is already exists: %w", errs.ErrAlreadyExists)
	} else {
		log.Printf("could not connect to database: %v\n", errs.ErrDatabaseConnection)

		return 0, fmt.Errorf("could not connect to database: %w", errs.ErrDatabaseConnection)
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
	log.Println("---redis id: ", idStr, "---")
	log.Printf("---redis error: %v\n", err)
	if err == redis.Nil {
		idStr = "0"
		err = mainDB.Set(ctx, "id", idStr, 0).Err()
		if err != nil {
			log.Fatalf("could not set id key: %v\n", err)
		}
	} else if err != nil {
		log.Printf("redis: %v", err)
		panic(err)
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
