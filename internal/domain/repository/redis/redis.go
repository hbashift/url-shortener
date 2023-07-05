package redis

import (
	"context"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

var ID uint64 = 0

type redisDb struct {
	ctx    context.Context
	client *redis.Client
}

type Config struct {
	Addr  string
	Pass  string
	DbNum int
}

// TODO передавать по указателю или по значению?

func (r *redisDb) GetUrl(shortUrl uint64) (string, error) {
	val, err := r.client.Get(r.ctx, strconv.FormatUint(shortUrl, 10)).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	}

	if err != nil {
		log.Printf("could not parse id: %v", err)
	}

	return val, nil
}

// TODO передавать по указателю или по значению?

func (r *redisDb) PostUrl(longUrl string) (uint64, error) {
	// TODO autoincrement id
	ID++
	err := r.client.Set(r.ctx, strconv.FormatUint(ID, 10), longUrl, 0).Err()
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func initRedis(cfg *Config) (*redis.Client, context.Context) {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       cfg.DbNum,
	}), context.Background()
}

func NewRedis(cfg *Config) repository.Repository {
	client, ctx := initRedis(cfg)

	return &redisDb{
		ctx:    ctx,
		client: client,
	}
}
