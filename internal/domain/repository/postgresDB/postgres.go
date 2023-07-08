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

// Config - config for PostgreSQL database
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB - creates new repository.Repository
func NewPostgresDB(cfg *Config) repository.Repository {
	db, err := initPostgresDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to PostgreSQL DB: %v", err)
	}

	return &postgresDb{db: db}
}

// initPostgresDB - creates new connection with PostgreSQL database and
// migrates model.Url struct into database
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

// isExists - checks if specified value exists in the database
func (p *postgresDb) isExists(columnName, value string) bool {
	var url model.Url
	r := p.db.Where(columnName+" = ?", value).Find(&url)
	if r.Error != nil {
		return true
	}

	return r.RowsAffected > 0
}

// GetUrl - get method for url from PostgreSQL database.
// If byLongUrl true - gets short url by long_url
// Else - gets by short_url
func (p *postgresDb) GetUrl(url *model.Url, byLongUrl bool) (string, error) {
	shortUrl := url.ShortUrl
	longUrl := url.LongUrl
	res := model.Url{}
	var err error

	if byLongUrl {
		err = p.db.Where("long_url = ?", longUrl).First(&res).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

				return "", fmt.Errorf("cannot get url: %w", errs.ErrNotFound)
			}

			return "", fmt.Errorf("db error: %w", err)
		}

		return res.ShortUrl, nil
	}

	err = p.db.Where("short_url = ?", shortUrl).First(&res).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return "", fmt.Errorf("cannot get url: %w", errs.ErrNotFound)
		}

		return "", fmt.Errorf("db error: %w", err)
	}

	return res.LongUrl, nil
}

// PostUrl -creates new record in PostgreSQL database.
// If long_url exists - returns errs.ErrLongUrlExists.
// If short_url exists - return errs.ErrShortUrlExists.
// Else - returns pgconn.PgError
func (p *postgresDb) PostUrl(url *model.Url) (string, error) {
	if p.isExists("long_url", url.LongUrl) {

		return "", fmt.Errorf("record already exists: %w", errs.ErrLongUrlExists)
	}

	err := p.db.Select("long_url", "short_url").Create(&url).Error
	UniqueViolationErr := &pgconn.PgError{Code: "23505"}

	if err != nil {
		if err != nil && errors.As(err, &UniqueViolationErr) {

			return "", fmt.Errorf("short url already exists: %w", errs.ErrShortUrlExists)
		}

		return "", fmt.Errorf("db error: %w", err)
	}

	return url.ShortUrl, nil
}
