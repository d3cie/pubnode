package db

import (
	"errors"

	"github.com/d3cie/pubnode/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	cfg := config.Get()

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "libsql",
		DSN:        cfg.DBDSN,
	}), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Tx() *DB {
	return &DB{db.Begin()}
}

func (DB) IsErrUniqueConstraintViolation(err error, cols []string) bool {
	panic("NOT IMPLEMENTED YET")
	return false
}

func (DB) IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
