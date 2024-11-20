package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model

	UserID    string
	ExpiresAt time.Time
	IpAddress *string
	UserAgent *string
}
