package entity

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}
