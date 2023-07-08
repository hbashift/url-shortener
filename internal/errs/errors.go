package errs

import "errors"

var (
	ErrInsertion          = errors.New("insertion failed")
	ErrNotFound           = errors.New("not found")
	ErrAlreadyExists      = errors.New("already exists")
	ErrDatabaseConnection = errors.New("could not connect to database")
	ErrDatabaseMigr       = errors.New("could not migrate structs to database")
	ErrLongUrlExists      = errors.New("such url already exists")
	ErrShortUrlExists     = errors.New("such short url already exists")
)
