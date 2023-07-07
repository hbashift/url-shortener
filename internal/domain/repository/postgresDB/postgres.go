package postgresDB

import (
	"errors"
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"github.com/hbashift/url-shortener/internal/errs"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type postgresDb struct {
	db *gorm.DB
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
	url := model.Url{}
	err := p.db.Take(&url, shortUrl).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("url not found: %v\n", err)

			return "", fmt.Errorf("cannot get url: %w", errs.ErrNotFound)
		}

		return "", fmt.Errorf("db error: %w", err)
	}

	return url.Url, nil
}

func (p *postgresDb) PostUrl(longUrl string) (uint64, error) {
	url := model.Url{Url: longUrl}
	err := p.db.Select("url").Create(&url).Error
	UniqueViolationErr := &pgconn.PgError{Code: "23505"}

	if err != nil {
		if err != nil && errors.As(err, &UniqueViolationErr) {
			log.Printf("could not insert new record: %v", gorm.ErrDuplicatedKey)

			return 0, fmt.Errorf("record already exists: %w", errs.ErrAlreadyExists)
		}

		return 0, fmt.Errorf("db error: %w", err)
	}

	return url.ID, nil
}

func NewPostgresDB(cfg *Config) repository.Repository {
	db, err := initPostgresDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to PostgreSQL DB: %v", err)
	}

	return &postgresDb{db: db}
}

func initPostgresDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s user=%s",
		cfg.Host, cfg.Port, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.Username)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if err == gorm.ErrInvalidDB {
			log.Printf("could not connect to database: %v\n", err)

			return nil, fmt.Errorf("could not connect to database: %w", err)
		}

		return nil, err
	}

	err = db.AutoMigrate(&model.Url{})
	if err != nil {
		log.Printf("could not migrate database: %v\n", err)

		return nil, fmt.Errorf("could not migrate database: %w", errs.ErrDatabaseMigr)
	}

	return db, nil
}
