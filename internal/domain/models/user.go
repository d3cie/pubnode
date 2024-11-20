package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            string
	Username      string `gorm:"unique"`
	Email         string `gorm:"unique"`
	AvatarUrl     *string
	Name          *string
	GithubID      *string `gorm:"unique"`
	AuthProviders []string
	PasswordHash  []byte
}
