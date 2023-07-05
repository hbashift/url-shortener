package postgres

import (
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // TODO use gorm
	"log"
)

type postgresDb struct {
	postgres *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (p *postgresDb) GetUrl(shortUrl uint64) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (p *postgresDb) PostUrl(longUrl string) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func NewPostgresDB(cfg Config) repository.Repository {
	db, err := InitPostgresDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to PostgreSQL DB: %v", err)
	}

	return &postgresDb{postgres: db}
}

func InitPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
