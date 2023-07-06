package postgres

import (
	"fmt"
	"github.com/hbashift/url-shortener/internal/domain/errs"
	"github.com/hbashift/url-shortener/internal/domain/repository"
	"github.com/hbashift/url-shortener/internal/domain/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var notFoundError errs.NotFound
var alreadyExistsError errs.AlreadyExists
var connectionError errs.DatabaseConnectionError
var migrationError errs.DatabaseMigrationError

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
	result := p.db.Take(&url, shortUrl)

	if result.Error != nil {
		log.Printf("url not found: %v\n", result.Error)
		notFoundError = fmt.Errorf("url not found: %w", result.Error)

		return "", notFoundError
	}

	return url.Url, nil
}

func (p *postgresDb) PostUrl(longUrl string) (uint64, error) {
	url := model.Url{Url: longUrl}
	result := p.db.Select("url").Create(&url)

	if result.Error != nil {
		log.Printf("could not insert new record: %v", result.Error)
		alreadyExistsError = fmt.Errorf("record already exists: %w", result.Error)

		return 0, alreadyExistsError
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
		log.Printf("could not connect to database: %v\n", err)
		connectionError = fmt.Errorf("could not connect to database: %w", err)

		return nil, connectionError
	}

	err = db.AutoMigrate(&model.Url{})
	if err != nil {
		log.Printf("could not migrate database: %v\n", err)
		migrationError = fmt.Errorf("could not migrate database: %w", err)

		return nil, migrationError
	}

	return db, nil
}
