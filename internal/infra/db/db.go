package db

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func New(dsn string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open("./data/local.sqlite"), &gorm.Config{
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
