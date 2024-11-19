package db

import (
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
