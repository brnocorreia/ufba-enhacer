package database

import (
	"fmt"

	"github.com/brnocorreia/ufba-enhacer/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST = "DB_HOST"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
)

func NewDBConnection() (*gorm.DB, error) {
	conf := configs.GetDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	return db, nil
}
